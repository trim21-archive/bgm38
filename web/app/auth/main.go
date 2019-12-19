package auth

import (
	"github.com/gin-gonic/gin"
)

//Part bind auth part to gin app
func Part(app *gin.Engine) {
	var router = app.Group("/auth")
	router.GET("/v1/bgmtv.tv/callback", callback)
	router.GET("/v1/bgmtv.tv/redirect", redirect)
}
