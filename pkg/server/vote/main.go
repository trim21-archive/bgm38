package vote

import (
	"github.com/gin-gonic/gin"
)

//Part bind vote routers to gin app
func Part(app *gin.Engine) {
	indexPart(app)
}

func indexPart(app *gin.Engine) {
	var router = app.Group("/vote")
	router.GET("/", index)
	router.GET("/create", createUI)
	router.POST("/create", create)
	router.GET("/json", json)
	router.GET("/svg/:id", svg)
}
