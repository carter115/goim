package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gmimo/comet/connection"
	"gmimo/comet/stub"
	"gmimo/common/constant"
	"gmimo/common/e"
	"gmimo/common/log"
	messageProto "gmimo/common/proto/message"
	_ "gmimo/common/swagger"
	"net/http"
	"time"
)

// @Summary 发送消息
// @Description <li>1. id: 消息唯一ID</li><li><b>2. srcId: 发送该消息的ID</b></li><li>3. dstId: 接收消息的ID</li><li><b>4. msgType: 消息类型(1聊天室 2单聊 3群聊 4广播 5组播)</b></li><li>5. content: 消息内容</li><li>6. resType: 资源类型(1图片 2声音 3视频)</li><li>6. resUrl: 资源URL</li><li>7. CreateTime: 消息创建时间</li><li>8. ReadTime: 读消息时间</li>
// @Tags 消息
// @Accept  json
// @Param message body swagger.Message true "消息体"
// @Success 200 {object} swagger.Response
// @Router /message/send [post]
func MessageSend(c *gin.Context) {
	var msg = messageProto.NewMessage()
	msg.CreateTime = time.Now().Unix()
	defer msg.Put() // 放回pool

	// 数据校验
	_ = c.ShouldBindBodyWith(&msg, binding.JSON)

	err := validate.Struct(msg)
	if err != nil || !constant.MessageTypeSet.Contain(msg.MsgType) {
		log.Warnc(c, e.GetMsg(e.INVALID_PARAMS))
		c.JSON(http.StatusOK, e.NewResp(e.INVALID_PARAMS, nil))
		return
	}

	// rpc调用logic的消息服务
	result, errCode := stub.CallSendMessage(c, msg)
	c.JSON(http.StatusOK, e.NewResp(errCode, result))
}

// 消息分发到websocket用户连接
func MessageDistribute() {
	for {
		select {
		case msg := <-stub.PushMessageChan:
			if msg.MsgType == constant.P2Room {
				// 房间消息
				conns := connection.LoadRoomWsConnection(msg.DstId)
				log.Debugf("[%s]房间的连接:%v, 消息ID:%s", msg.DstId, conns, msg.Id)
				for _, conn := range conns {
					go conn.WriteMessage([]byte(msg.Content))
				}
			} else if msg.MsgType == constant.P2Peer {
				// 个人消息
				conn := connection.LoadUserWsConnection(msg.DstId)
				go conn.WriteMessage([]byte(msg.Content))
			}
		}
	}
}
