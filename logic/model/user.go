package model

// 设置用户所在房间ID
func UserSetRoom(key string, mid string) (err error) {
	_, err = RedisPool.Set(key, mid, defaultExpire).Result()
	return err
}

// 查询用户所在房间ID
func UserGetRoom(key string) (mid string, err error) {
	return RedisPool.Get(key).Result()
}

// 删除用户所在房间ID
func DelUserKey(key string) (err error) {
	return RedisPool.Del(key).Err()
}
