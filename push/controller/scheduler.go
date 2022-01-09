package controller

import (
	"context"
	"gmimo/common/middleware"
	messageProto "gmimo/common/proto/message"
	"gmimo/common/util"
	"gmimo/push/stub"
)

// [Push服务]从MQ拿取消息，转存到Mongodb，并且推送消息给[Comet服务]

var Scheduler *scheduler

func InitScheduler() {
	if Scheduler == nil {
		Scheduler = &scheduler{
			inChan:  make(chan *messageProto.Request, 1000),
			outChan: make(chan *messageProto.Request, 1000),
		}
		go Scheduler.Dispatch() //启动分发消息
	}
	return
}

type scheduler struct {
	inChan  chan *messageProto.Request
	outChan chan *messageProto.Request
}

// 推送给comet服务,转存Mongo
func (s *scheduler) Dispatch() {
	for {
		select {
		case msg := <-s.inChan:
			// TODO rpc调用comet
			ctx := middleware.SetReqidToRpcContext(util.GetUUID(), context.Background())
			go stub.CallSendMessage(ctx, msg)
			// 消息存储到DB
			go saveMessage(msg)
		}
	}
}

// 提供给其它协程来提交消息
func (s *scheduler) Push(msg *messageProto.Request) {
	s.inChan <- msg
}
