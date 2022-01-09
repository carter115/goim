package connection

// 通过Redis的发布/订阅，来同步多个comet的用户连接状态

import (
	"encoding/json"
	"gmimo/comet/config"
	"gmimo/comet/model"
	"gmimo/common/log"
	"os"
	"time"
)

var (
	defaultChannelName = "userConn"
	JoinRoom           = "JoinRoom"
	LeaveRoom          = "LeaveRoom"
)

// 定义通知内容: 房间ID,用户ID,连接ID
type UserConnEvent struct {
	Uid    string `json:"uid"`
	ConnId string `json:"connId"`
	Action string `json:"action"`
}

// 发布
func Publish(event UserConnEvent) error {
	bs, _ := json.Marshal(event)
	return model.RedisPool.Publish(defaultChannelName, bs).Err()
}

// 订阅
func Subscribe() {
	if err := model.InitRedisClient(config.Redis); err != nil {
		log.Error("failed to connect redis:", err)
		os.Exit(-1)
	}
	pubsub := model.RedisPool.Subscribe(defaultChannelName)
	defer pubsub.Close()

	log.Info("connection subscribe is running")
	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			log.Error("pubsub recv error:", err)
			time.Sleep(time.Second)
			continue
		}
		event := UserConnEvent{}
		json.Unmarshal([]byte(msg.Payload), &event)
		// 移除本地的用户连接状态
		log.Infof("recv event: %+v", event)
		val, ok := WsConnPoolMap.Load(event.Uid)
		if ok && event.ConnId == val && event.Action == LeaveRoom {
			WsConnPoolMap.Delete(event.Uid)
			log.Info("移除本地缓存的用户连接", event)
		}
	}
}
