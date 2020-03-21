package cron

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"bgm38/pkg/db"
)

const (
	typeMusic = "Music"
	chunkSize = 5000
)

var m = sync.Mutex{}
var maxMapID = 0
var blockList = []string{"角色出演", "片头曲", "片尾曲", "其他", "画集", "原声集"}

func reCalculateMap() {
	var err error
	var maxSubject db.Subject
	var minSubject db.Subject
	maxMapID = 0

	err = db.MysqlX.Get(&maxSubject, `select * from subject order by id desc limit 1`)
	check(err)
	err = db.MysqlX.Get(&minSubject, `select * from subject order by id limit 1`)
	check(err)

	tx, err := db.MysqlX.Beginx()
	if err != nil {
		logrus.Errorln(err)
		return
	}

	if _, err := tx.Exec(`UPDATE relation SET removed = 0 WHERE true`); err != nil {
		logrus.Errorln(err)
		return
	}
	preRemove(tx, minSubject.ID, maxSubject.ID)
	err = tx.Commit()
	if err != nil {
		logrus.Errorln(err)
		return
	}

	tx, err = db.MysqlX.Beginx()
	if err != nil {
		logrus.Errorln(err)
		return
	}

	err = firstRun(tx, minSubject.ID, maxSubject.ID)
	if err != nil {
		_ = tx.Rollback()
		logrus.Errorln(err)
		return
	}
	err = tx.Commit()
	if err != nil {
		logrus.Errorln(err)
		return
	}

}

func removeNodes(tx *sqlx.Tx, nodeIDs ...int) {
	query, args, err := sqlx.In(`UPDATE subject SET locked = 1 WHERE id IN (?)`, nodeIDs)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}

	query = tx.Rebind(query)
	tx.MustExec(query, args...)

	query, args, err = sqlx.In(`UPDATE relation SET removed = 1 WHERE target in (?) or source in (?)`, nodeIDs, nodeIDs)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}
	query = tx.Rebind(query)
	tx.MustExec(query, args...)
}

func relationsNeedToRemove(tx *sqlx.Tx, m map[int]int) {
	for id1, id2 := range m {
		tx.MustExec(`UPDATE relation SET removed = 1 
WHERE (target = ? AND source = ?) OR (target = ? AND source = ?)`, id1, id2, id2, id1)
	}
}

func preRemove(tx *sqlx.Tx, subjectStart int, subjectEnd int) {
	println("pre remove")

	removeNodes(tx, 91493, 102098, 228714, 231982, 932, 84944, 78546)

	query, args, err := sqlx.In(`update relation set removed=1 where relation in (?)`, blockList)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}
	query = tx.Rebind(query)
	tx.MustExec(query, args...)

	relationsNeedToRemove(tx, map[int]int{
		91493:  8,
		8108:   35866,
		446:    123207,
		123207: 27466,
		123217: 4294,
	})

	tx.MustExec(`update subject set locked=1 where subject_type=?`, typeMusic)
	var idToRemove = make(map[int]bool)
	var subjects []db.Subject
	err = tx.Select(&subjects, `SELECT * FROM subject WHERE locked = ?`, 1)
	check(err)

	for _, s := range subjects {
		idToRemove[s.ID] = true
	}

	var idToRM []int
	for key := range idToRemove {
		idToRM = append(idToRM, key)
	}

	err = chunkIterInt(idToRM, func(s []int) error {
		return execIn(tx, `UPDATE relation SET removed=1 where source IN (?) OR target IN (?)`, s, s)
	})

	if err != nil {
		logrus.Errorln(err)
	}
	for i := subjectStart; i <= subjectEnd; i += chunkSize {
		relationIDNeedToRemove := make(map[string]bool)
		sourceToTarget := make(map[int]map[int]bool)
		var relations []db.Relation
		err := tx.Select(&relations, `SELECT * FROM relation WHERE (( source >= ? AND source < ? ) OR ( target >= ? AND target < ? ) ) AND removed = ?`,
			i, i+chunkSize, i, i+chunkSize, 0)
		check(err)
		for _, rel := range relations {
			if sourceToTarget[rel.Source] == nil {
				sourceToTarget[rel.Source] = make(map[int]bool)
			}
			sourceToTarget[rel.Source][rel.Target] = true
		}

		for _, rel := range relations {
			if subMap, ok := sourceToTarget[rel.Target]; ok {
				if _, ok := subMap[rel.Source]; ok {
					continue
				}
			}
			relationIDNeedToRemove[rel.ID] = true
		}

		var ids = make([]string, 0, len(relationIDNeedToRemove))
		for key := range relationIDNeedToRemove {
			ids = append(ids, key)
		}
		if len(ids) != 0 {
			_ = execIn(tx, `update relation set removed = 1 where id IN (?)`, ids)
		}
	}
	println("finish pre remove")
}

func getRelationsFromDB(tx *sqlx.Tx, subjectStart, subjectEnd int) (map[int]map[string]*db.Relation, int, error) {

	edgeCount := 0
	var relationFromID = make(map[int]map[string]*db.Relation)
	for i := subjectStart; i < subjectEnd; i += chunkSize {
		var edges = make([]*db.Relation, 0, 5000)

		err := tx.Select(&edges, `select * from relation where (source >= ?) AND (source < ?) AND (removed = ?)`, i, i+chunkSize, 0)

		if err != nil {
			return nil, 0, err
		}

		for _, edge := range edges {
			if edge.Removed == 1 {
				continue
			}
			edge.Map = 0
			edgeCount++
			if relationFromID[edge.Source] == nil {
				relationFromID[edge.Source] = make(map[string]*db.Relation)
			}

			if relationFromID[edge.Target] == nil {
				relationFromID[edge.Target] = make(map[string]*db.Relation)
			}

			relationFromID[edge.Source][edge.ID] = edge
			relationFromID[edge.Target][edge.ID] = edge
		}
	}

	return relationFromID, edgeCount, nil
}
func getSubjectsFromDB(tx *sqlx.Tx, subjectStart int, subjectEnd int) (map[int]*db.Subject, error) {
	var subjects = make(map[int]*db.Subject)
	for i := subjectStart; i < subjectEnd; i += chunkSize {
		var s = make([]*db.Subject, 0, 5000)

		err := tx.Select(&s, `select * from subject where id >= ? AND id < ? AND locked = ? AND subject_type != ?`, i, i+chunkSize, 0, typeMusic)
		if err != nil {
			return nil, err
		}

		for _, subject := range s {
			if subject.SubjectType == typeMusic || subject.Locked != 0 {
				logrus.Errorf("subject error %v\n", subject.ID)
				continue
			}
			subject.Map = 0

			subjects[subject.ID] = subject
		}
		s = nil
	}
	return subjects, nil
}

func firstRun(tx *sqlx.Tx, subjectStart int, subjectEnd int) error {
	logrus.Debugf("build relation map with start id %d and end id %d", subjectStart, subjectEnd)
	var doneIDs = make(map[int]bool, subjectEnd-subjectStart)
	subjects, err := getSubjectsFromDB(tx, subjectStart, subjectEnd)
	if err != nil {
		return err
	}
	for key, value := range subjects {
		if key != value.ID {
			panic(fmt.Sprintf("value of key %d is not %v", key, value))
		}

	}
	logrus.Infof("total subject %d", len(subjects))

	relationFromID, edgeCount, err := getRelationsFromDB(tx, subjectStart, subjectEnd)

	if err != nil {
		return err
	}

	if m, ok := relationFromID[8108]; !ok {
		logrus.Fatalln("err")
	} else {
		for id, edge := range m {
			if edge.Removed == 1 {
				fmt.Println(id, edge)
			}
		}
	}

	logrus.Infof("total %d edges", edgeCount)
	logrus.Debugf("id 8 get %d edges", len(relationFromID[8]))

	count := 0
	var dealWithNode func(int)
	dealWithNode = func(sourceID int) {
		if _, ok := doneIDs[sourceID]; ok {
			return
		}
		count++
		var edges = relationFromID[sourceID]
		var mapID = 0
		for _, edge := range edges {
			if edge.Map != 0 {
				mapID = edge.Map
				break
			}
			if subjects[edge.Target].Map != 0 {
				mapID = subjects[edge.Target].Map
				break
			}
			if subjects[edge.Source].Map != 0 {
				mapID = subjects[edge.Source].Map
				break
			}
		}

		if mapID == 0 {
			m.Lock()
			maxMapID++
			mapID = maxMapID
			m.Unlock()
		}
		for k := range edges {
			edges[k].Map = mapID
		}
		subjects[sourceID].Map = mapID
		doneIDs[sourceID] = true
		for _, edge := range edges {
			dealWithNode(edge.Target)
		}
	}

	logrus.Debugf("now iter %d subjects", len(subjects))

	for id := range subjects {
		dealWithNode(id)
	}

	logrus.Debugf("called %d times", count)
	logrus.Debugf("done %d ids", len(doneIDs))

	var subjectMaps = make(map[int][]int)
	var relationMaps = make(map[int][]string)
	for _, subject := range subjects {
		subjectMaps[subject.Map] = append(subjectMaps[subject.Map], subject.ID)
	}

	for _, edges := range relationFromID {
		for _, edge := range edges {
			relationMaps[edge.Map] = append(relationMaps[edge.Map], edge.ID)
		}
	}

	logrus.Infof("got %d map", len(subjectMaps))
	for key, ids := range subjectMaps {
		logrus.Debugln(key)
		err = chunkIterInt(ids, func(s []int) error {
			return execIn(tx, `update subject set map=? where id in (?)`, key, s)
		})
		if err != nil {
			return nil
		}
	}

	for key, ids := range relationMaps {
		logrus.Debugln(key)
		err := chunkIterStr(ids, func(s []string) error {
			return execIn(tx, `update relation set map = ? where id IN (?)`, key, s)
		})
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(`update subject set locked=? where map=?`, 1, 0)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`update relation set removed=? where map=?`, 1, 0)
	if err != nil {
		return err
	}
	logrus.Infoln("finish save to db")
	return nil
}

func chunkIterInt(s []int, f func([]int) error) error {
	var err error
	l := len(s)
	for i := 0; i < l; i += chunkSize {
		err = f(s[i:min(i+chunkSize, l)])
		if err != nil {
			return err
		}
	}
	return err
}

func chunkIterStr(s []string, f func([]string) error) error {
	var err error
	l := len(s)
	for i := 0; i < l; i += chunkSize {
		err = f(s[i:min(i+chunkSize, l)])
		if err != nil {
			return err
		}
	}
	return err
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func check(err error) {
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}

func execIn(tx *sqlx.Tx, query string, args ...interface{}) error {
	q, a, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}
	q = tx.Rebind(q)
	_, err = tx.Exec(q, a...)
	return err
}
