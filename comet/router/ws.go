package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gmimo/comet/connection"
	"gmimo/comet/stub"
	"gmimo/common/e"
	"gmimo/common/log"
	roomProto "gmimo/common/proto/room"
	"net/http"
	"time"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// @Summary 发起websocket连接
// @Tags websocket
// @param uid query string true "用户ID"
// @param mid query string true "房间ID"
// @Success 200
// @Router /ws/connect [get]
func WsConnect(c *gin.Context) {
	var (
		err error
	)
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Errorc(c, "failed to websocket connection:", err.Error())
		c.JSON(http.StatusOK, e.NewResp(e.ERR_WS_ERROR, nil))
		return
	}
	uid := c.Query("uid")
	mid := c.Query("mid")
	if len(uid) == 0 || len(mid) == 0 {
		c.JSON(http.StatusBadRequest, e.NewResp(e.INVALID_PARAMS, nil))
		return
	}

	// rpc调用logic加入房间
	req := &roomProto.Request{Mid: mid, Uid: uid}
	if _, errCode := stub.CallRoomJoin(c, req); errCode != e.SUCCESS {
		c.JSON(http.StatusOK, e.NewResp(errCode, nil))
		return
	}

	// 得到websocket的长链接之后,就可以对客户端传递的数据进行操作
	wsConn, _ := connection.InitWsConnection(uid, mid, conn)
	if wsConn != nil {
		connection.WsConnPoolMap.Store(wsConn.Uid, wsConn)
	}
}

// @Summary 关闭websocket连接
// @Tags websocket
// @param uid query string true "用户ID"
// @Success 200
// @Router /ws/close [get]
func WsClose(c *gin.Context) {
	uid := c.Query("uid")
	if len(uid) == 0 {
		c.JSON(http.StatusBadRequest, e.NewResp(e.INVALID_PARAMS, nil))
		return
	}

	if conn := connection.LoadUserWsConnection(uid); conn != nil {
		conn.Close()
	}
	c.JSON(http.StatusOK, e.NewResp(e.SUCCESS, nil))
}

// @Summary 列出websocket连接用户
// @Tags websocket
// @Success 200
// @Router /ws/list [get]
func WsConnectList(c *gin.Context) {
	list := connection.ConnPoolMapList()
	c.JSON(http.StatusOK, e.NewResp(e.SUCCESS, list))
}
