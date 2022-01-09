package controller

import (
	"gmimo/common/constant"
	"gmimo/common/log"
	messageProto "gmimo/common/proto/message"
	"gmimo/push/model"
)

func saveMessage(msg *messageProto.Request) {
	switch msg.MsgType {
	case constant.P2Room:
		saveToRedis(msg) // 存储到Redis,设置TTL
	case constant.P2Peer, constant.P2Group, constant.Broadcast:
		saveToMongo(msg) // 存储到Mongo
	}
}

// 存储到redis
func saveToRedis(msg *messageProto.Request) {
	if err := model.SaveToRedis(msg); err != nil {
		log.Error("save to redis error:", err)
	}
}

// 存储到mongo
func saveToMongo(msg *messageProto.Request) {
	if err := model.SaveToMongo(msg); err != nil {
		log.Error("save to mongo error:", err)
	}
}
