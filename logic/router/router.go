package router

import (
	"github.com/gin-gonic/gin"
	"gmimo/common/middleware"
	"gmimo/logic/config"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	if !config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// 加载中间件
	engine.Use(middleware.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.SetRequestId()) // 添加请求ID

	engine.GET("", func(c *gin.Context) {
		c.String(200, "home page")
	})

	// 路由分组
	//room := engine.Group("/room")
	//{
	//	room.GET("/join", RoomJoin)
	//	room.GET("/leave", RoomLeave)
	//	room.GET("/count", RoomCount)
	//	room.GET("/member", RoomMember)
	//}

	//message := engine.Group("/message")
	//{
	//	message.GET("/push", )
	//}

	return engine
}
