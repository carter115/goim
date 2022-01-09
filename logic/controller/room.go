package controller

import (
	"fmt"
	"gmimo/common/log"
	"gmimo/logic/model"
)

// 处理房间的业务: 加入，离开，用户列表，用户数

type Room struct {
	Id string
}

// 加入房间
func (m *Room) Join(uid string) (err error) {
	// 用户先离开旧的房间
	if mid, err := model.UserGetRoom(userKey(uid)); err == nil {
		model.RoomDelUser(roomKey(mid), uid)
	}

	// 设置用户所在房间
	if err = model.UserSetRoom(userKey(uid), m.Id); err != nil {
		log.Error("设置用户所在房间ID失败")
		return
	}

	// 房间ID添加该用户
	if err = model.RoomAddUser(roomKey(m.Id), uid); err != nil {
		log.Warn("添加房间用户失败")
		return
	}

	return
}

// 离开房间
func (m *Room) Leave(uid string) (err error) {
	// 删除用户所有的房间, 忽略删除失败
	if err = model.DelUserKey(userKey(uid)); err != nil {
		log.Warn("删除用户所在房间ID失败")
	}

	// 房间ID移除该用户
	if err = model.RoomDelUser(roomKey(m.Id), uid); err != nil {
		log.Warn("移除房间用户失败")
		return
	}

	return
}

// 房间用户数
func (m *Room) Count() int64 {
	n, err := model.RoomUserCount(roomKey(m.Id))
	if err != nil {
		log.Warn("获取房间用户数失败:" + err.Error())
	}
	return n
}

// 房间用户列表
func (m *Room) Member() ([]string, error) {
	return model.RoomMember(roomKey(m.Id))
}

// -----------------定义Redis Key-------------------

// 定义在Redis中的Key
func userKey(uid string) string {
	return fmt.Sprintf("user:%s", uid)
}

func roomKey(mid string) string {
	return fmt.Sprintf("room:%s", mid)
}
