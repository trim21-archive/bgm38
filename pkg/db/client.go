package db

import (
	"bgm38/config"
	"github.com/go-redis/redis"
)

// Redis redis client
var Redis = redis.NewClient(&redis.Options{
	Addr:     config.RedisAddr,
	Password: config.RedisPassword, // no password set
	// Mysql:       0,            // use default Mysql
})