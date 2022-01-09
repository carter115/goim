package model

import (
	"fmt"
	"gmimo/common/log"
	"gmimo/logic/config"
	"testing"
)

// 项目初始化
func Init() (err error) {
	if err = config.InitConfig(); err != nil {
		fmt.Println("初始化配置模块失败:", err)
		return
	}
	fmt.Println("成功加载配置模块", config.String())

	if err = log.InitLogger(config.Logger, config.App.Name); err != nil {
		fmt.Println("初始化日志模块失败:", err)
		return
	}
	fmt.Println("成功加载日志器")

	if err = InitRedisClient(config.Redis); err != nil {
		log.Error("Redis server 连接失败")
		return
	}
	log.Info("Redis server 连接成功")

	return
}

func TestModel(t *testing.T) {
	var err error
	if err = Init(); err != nil {
		t.Fatal(err)
	}

	// 用户
	userKey := "user:Jackaa"
	//if err = UserSetRoom(userKey, "abc"); err != nil {
	//	t.Log(err)
	//}

	mid, err := UserGetRoom(userKey)
	if err != nil {
		t.Fatalf("%q", err)
	}
	t.Logf("%q,%q", userKey, mid)

	// 房间
	//roomKey := "room:abc"
	//s, err := RoomMember(roomKey)
	//n, _ := RoomUserCount(roomKey)
	//t.Log("房间用户列表和用户数:", s, n, err)

	//if err = DelUserKey(userKey); err != nil {
	//	t.Error(err)
	//}

}
