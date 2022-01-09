package constant

import "gmimo/common/util"

// byte size
const (
	B  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

// 请求ID
const (
	REQID = "reqid"
)

//消息体类型: 1聊天室 2单聊 3群聊 4广播 5组播
const (
	P2Room    = "P2Room"
	P2Peer    = "P2Peer"
	P2Group   = "P2Group"
	Broadcast = "Broadcast"
	Multicast = "Multicast"
)

var MessageTypeSet = util.Set{}
var ResTypeSet = util.Set{}

//资源类型: 1图片 2声音 3视频
const (
	Pic   = "Pic"
	Sound = "Sound"
	Video = "Video"
)

func init() {
	MessageTypeSet.Push(P2Room, P2Peer, P2Group, Broadcast, Multicast)
	ResTypeSet.Push(Pic, Sound, Video)
}
