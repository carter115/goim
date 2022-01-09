package mimo_srv_message

import (
	"gmimo/common/util"
	"sync"
)

// 从池中，获取一条空消息
func NewMessage() *Request {
	buf := cachePool.Get() // 获取
	return buf.(*Request)
}

func (m *Request) Put() *Request {
	m = &Request{Id: util.GetUUID()}
	cachePool.Put(m) // 放回pool
	return m
}

// 消息缓冲池
var cachePool = sync.Pool{
	New: func() interface{} {
		return &Request{Id: util.GetUUID()}
	},
}
