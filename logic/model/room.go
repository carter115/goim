package model

// 房间添加用户
func RoomAddUser(key string, uid string) error {
	return RedisPool.SAdd(key, uid).Err()
}

// 房间移除用户
func RoomDelUser(key string, uid string) error {
	return RedisPool.SRem(key, uid).Err()
}

// 房间用户数
func RoomUserCount(key string) (int64, error) {
	return RedisPool.SCard(key).Result()
}

// 房间用户列表
func RoomMember(key string) ([]string, error) {
	return RedisPool.SMembers(key).Result()
}
