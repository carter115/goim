package main

import (
	"encoding/json"
	"gmimo/common/log"
	"gmimo/common/util"
)

type Message struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	Content   string `json:"content,omitempty"`
	Target    string `json:"target"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

// 基础的消息
func newMessage() (msg Message) {
	msg.Id = util.GetUUID()
	return
}

// 心跳消息体
func NewHearbeatMessage() Message {
	msg := newMessage()
	msg.Type = "HeartBeat"
	return msg
}

// 房间消息体
func NewRoomMessage(rid string, content string) Message {
	msg := newMessage()
	msg.Type = "P2Room"
	msg.Target = rid
	msg.Content = content
	return msg
}

// 加入房间消息体
func NewJoinRoomMessage(rid string) Message {
	msg := newMessage()
	msg.Type = "JoinRoom"
	msg.Target = rid
	return msg
}

// 离开房间消息体
func NewLeaveRoomMessage(rid string) Message {
	msg := newMessage()
	msg.Type = "LeaveRoom"
	msg.Target = rid
	return msg
}

// 反序列化消息体对象
func MessagePack(data []byte) (Message, error) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Warnf("marshal error: %s, message: %+v", err, msg)
	}
	msg.checkId()
	return msg, err
}

// 序列化消息体对象
func MessageUnPack(msg Message) ([]byte, error) {
	msg.checkId()
	bs, err := json.Marshal(msg)
	if err != nil {
		log.Error("marshal message error:", err, msg)
	}
	return bs, err
}

// 检查消息ID
func (msg *Message) checkId() {
	if len(msg.Id) != 32 {
		log.Warnf("message id length must be 32: %+v", msg)
	}
}
