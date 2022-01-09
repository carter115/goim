package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gmimo/comet/config"
	"gmimo/comet/connection"
	_ "gmimo/comet/docs"
	"gmimo/common/e"
	"gmimo/common/log"
	"gmimo/common/middleware"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	if config.App.Debug {
		config.App.JwtExpire = time.Hour * 24 * 30 // 测试环境，设定1个月Token过期时间
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 加载中间件
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recovery())
	engine.Use(middleware.SetRequestId()) // 添加请求ID

	// -------------------非业务路由-------------------
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/", Home)
	engine.GET("/hello", Hello)

	// -------------------业务路由-------------------

	// 生成token
	engine.GET("/auth", GetAuth)

	// 消息模块
	message := engine.Group("/message")
	{
		message.POST("/send", MessageSend)
	}
	go MessageDistribute() // 通过websocket连接，把消息分发给用户

	// 房间模块
	room := engine.Group("/room")
	room.Use(middleware.Jwt())
	{
		room.POST("/join", RoomJoin)
		room.POST("/leave", RoomLeave)
		room.GET("/member/:mid", RoomUserList) // TODO 没实现

	}

	// websocket模块
	ws := engine.Group("/ws")
	//ws.Use(middleware.Jwt())
	{
		ws.GET("/connect", WsConnect)
		ws.GET("/list", WsConnectList)
		ws.GET("/close", WsClose)
	}
	go connection.Subscribe() //启动订阅服务，用来维护多个comet的用户连接信息同步

	//// 用户模块
	//user := engine.Group("/user")
	//user.Use(middleware.Jwt()) // 添加token认证中间件
	//{
	//	user.GET("", User)
	//	user.POST("/update", UserUpdate)
	//}

	// 处理错误的请求: 不存在的路由或方法
	engine.NoRoute(func(c *gin.Context) {
		log.Warnc(c, e.GetMsg(e.NOT_ROUTE_METHOD))
		c.JSON(http.StatusNotFound, e.NewResp(e.NOT_ROUTE_METHOD, nil))
	})

	return engine
}

// 用于数据校验
var validate = validator.New()
