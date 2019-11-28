package vote

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Person struct {
	ID string `uri:"id" binding:"required"`
	//Name string `uri:"name" binding:"required"`
}

const f = `
<svg xmlns="http://www.w3.org/2000/svg" version="1.1">
  <circle cx="100" cy="50" r="40" stroke="black"
  stroke-width="2" fill="red" />
</svg>
`

func svg(ctx *gin.Context) {
	var person Person
	if err := ctx.ShouldBindUri(&person); err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{"msg": err})
		return
	}
	//ctx.JSON(200, gin.H{"id": person.ID})
	ctx.Render(200, SVG{Data: f})
	//writeContentType(ctx.Writer, svgContentType)
	//ctx.Writer.WriteString(f)

	//ctx.HTML(200, "vote/index.tmpl", map[string]string{
	//	"Title":   "hello Vote Function",
	//	"Message": "hello world",
	//})

}
