package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gmimo/common/log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type WsConnectionClient struct {
	Conn      *websocket.Conn // websocket 连接
	Uid       string          // 用户ID
	RecvCount int64           // 用于统计收到消息数量
	isClosed  bool
}

// 创建连接
func NewWsConnectionClient(wg *sync.WaitGroup, users []*WsConnectionClient, idx int) {
	// 增加休眠时间
	time.Sleep(time.Duration(rand.Intn(waitN)) * time.Millisecond)
	defer wg.Done()
	u := fmt.Sprintf("%s%d", preUid, idx)
	url := fmt.Sprintf("%s?uid=%s&token=%s&deviceId=%s&terminal=%s&version=%s",
		wsAddr, u, GetToken(), GetDeviceId(), GetTerminal(), GetVersion())
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Error("new websocket connection error:", err)
		return
	}
	cli := &WsConnectionClient{Conn: conn, Uid: u}

	// 连接后，开始循环读消息
	go cli.ReadLoop()
	log.Debugf("[%s] is connected.", cli.Uid)
	users[idx] = cli
}

// read 协程
func (w *WsConnectionClient) ReadLoop() {
	for {
		_, p, err := w.Conn.ReadMessage()
		if err != nil {
			//log.Error("read message error:", err)
			go w.Close() // 关闭连接，退出
			return
		}

		log.Debugf("[%s] recv message: %s", w.Uid, string(p))

		msg, err := MessagePack(p)
		if err != nil {
			continue
		}

		// 累计有效的接收消息数量
		if msg.Type != "Ack" && msg.Type != "Error" {
			atomic.AddInt64(&w.RecvCount, 1)
		}
	}
}

// write协程
func (w *WsConnectionClient) WriteLoop(msg Message) {
	for {
		// 房间消息体长度为0，则不发送
		if msg.Type == "P2Room" && len(msg.Content) == 0 {
			continue
		}

		w.writeMessage(msg)
	}
}

// 用户发送一条消息
func (w *WsConnectionClient) writeMessage(msg Message) {
	// panic recovery
	defer func() {
		if e := recover(); e != nil {
			log.Error("[%v] panic:", w, e)
			return
		}
	}()

	bytes, err := MessageUnPack(msg)
	if err != nil {
		return
	}

	// 并发发送消息，减少并发压力
	if w == nil {
		return
	}
	time.Sleep(time.Duration(waitTime+rand.Intn(waitN)) * time.Millisecond)
	if err := w.Conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
		//log.Error("write message error:", err)
		return
	}

	// 写一条消息成功
	log.Debugf("[%s] send message : %+v", w.Uid, msg)
	return
}

func (w *WsConnectionClient) Close() {
	if w.Conn != nil && !w.isClosed {
		w.isClosed = true
		log.Warnf("Closing connection: %+v", w)
		w.LeaveRoom()  // 离开房间
		w.Conn.Close() // 可重复关闭
	}
}

//// 发送房间消息(先加房间，再发消息)
//func (w *WsConnectionClient) SendRoomMessage(rid string, content string) {
//	w.JoinRoom()
//	msg := NewRoomMessage(rid, content)
//	w.writeMessage(msg)
//}

// 调用进入房间指令
func (w *WsConnectionClient) JoinRoom() {
	msg := NewJoinRoomMessage(roomId)
	w.writeMessage(msg)
}

// 调用离开房间指令
func (w *WsConnectionClient) LeaveRoom() {
	msg := NewLeaveRoomMessage(roomId)
	w.writeMessage(msg)
}

// 保持连接心跳
func (w *WsConnectionClient) Keeplive() {
	for {
		msg := NewHearbeatMessage()
		w.writeMessage(msg)
	}
}
