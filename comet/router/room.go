package router

import (
	"github.com/gin-gonic/gin"
	"gmimo/comet/stub"
	"gmimo/common/e"
	"gmimo/common/log"
	roomProto "gmimo/common/proto/room"
)

// @Summary 用户加入房间
// @Tags 房间管理
// @Param token header string true "Token"
// @Param mid query string true "房间ID"
// @Param uid query string true "用户ID"
// @Success 200
// @Router /room/join [post]
func RoomJoin(c *gin.Context) {
	mid := c.Query("mid")
	uid := c.Query("uid")

	// TODO 数据校验
	req := &roomProto.Request{Mid: mid, Uid: uid}
	// rpc调用logic加入房间
	result, errCode := stub.CallRoomJoin(c, req)
	c.JSON(200, e.NewResp(errCode, result))
}

// @Summary 用户离开房间
// @Tags 房间管理
// @Param token header string true "Token"
// @Param mid query string true "房间ID"
// @Param uid query string true "用户ID"
// @Success 200
// @Router /room/leave [post]
func RoomLeave(c *gin.Context) {
	mid := c.Query("mid")
	uid := c.Query("uid")

	// TODO 数据校验
	req := &roomProto.Request{Mid: mid, Uid: uid}
	stub.CallRoomLeave(c, req)
	c.JSON(200, e.NewResp(e.SUCCESS, nil))
}

// @Summary 该房间用户列表
// @Tags 房间管理
// @Param token header string true "Token"
// @Param mid path string true "房间ID"
// @Success 200
// @Router /room/member/{mid} [get]
func RoomUserList(c *gin.Context) {
	mid := c.Param("mid")
	req := &roomProto.Request{Mid: mid}
	rsp, code := stub.CallRoomMember(c, req)
	log.Debugcf(c, "room [%s] member response: %+v", mid, rsp)

	result := rsp.(*roomProto.RespMember)
	c.JSON(200, e.NewResp(code, result.Member))
}
