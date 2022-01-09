package model

import (
	"context"
	"encoding/json"
	"fmt"
	messageProto "gmimo/common/proto/message"
)

// redis
func SaveToRedis(msg *messageProto.Request) error {
	key := fmt.Sprintf("message:%s:%s:%s", msg.DstId, msg.SrcId, msg.Id)
	value, _ := json.Marshal(msg)
	return RedisPool.Set(key, value, defaultExpire).Err()
}

// mongo
func SaveToMongo(msg *messageProto.Request) error {
	record := Record{
		Id:         msg.Id,
		SrcId:      msg.SrcId,
		DstId:      msg.DstId,
		MsgType:    msg.MsgType,
		Content:    msg.Content,
		ResType:    msg.ResType,
		ResUrl:     msg.ResUrl,
		CreateTime: msg.CreateTime,
	}
	_, err := MongoCollection.InsertOne(context.Background(), record)
	return err
}

type Record struct {
	Id         string `bson:"_id"`
	SrcId      string `bson:"srcId"`
	DstId      string `bson:"dstId"`
	MsgType    string `bson:"msgType"`
	Content    string `bson:"content"`
	ResType    string `bson:"resType"`
	ResUrl     string `bson:"resUrl"`
	CreateTime int64  `bson:"createTime"`
	ReadTime   int64  `bson:"readTime"`
}
