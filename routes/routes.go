package routes

import "github.com/gin-gonic/gin"

func Routes(server *gin.Engine) {
	server.GET("/short", getUrls)
	server.POST("/short", addUrl)
	server.GET("/short/:short_code", getUrl)

}
