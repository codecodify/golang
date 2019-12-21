package initRouter

import (
	"github.com/gin-gonic/gin"
	"wechat/handlers"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	}else{
		router.LoadHTMLGlob("./templates/*")
	}
	router.GET("/", handlers.IndexHandler)
	router.GET("/api", handlers.ApiHandler)

	// 静态文件
	router.Static("/static", "./static")

	return router
}
