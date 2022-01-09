package e

var errMsgMap = map[int]string{
	SUCCESS:          "ok",
	ERROR:            "error",
	NOT_ROUTE_METHOD: "请求路径或方法不存在",
	INVALID_PARAMS:   "请求参数错误",

	ERR_RPC_ERROR:   "RPC错误",
	ERR_RPC_TIMEOUT: "RPC调用超时",

	ERR_WS_ERROR: "websocket错误",

	ERR_AUTH:            "Token错误",
	ERR_AUTH_TOKEN:      "Token生成失败",
	ERR_AUTH_TOKEN_FAIL: "Token校验失败",
	ERR_USER_PWD_FAIL:   "用户或密码错误",

	ERR_ROOM_JOIN:   "加入房间失败",
	ERR_ROOM_MEMBER: "获取房间用户失败",

	ERR_MESSAGE_SEND_FAIL: "发送消息失败",

	ERR_KAFKA_PRODUCER: "kafka生产者创建失败",
	ERR_KAFKA_CONSUMER: "kafka消费者创建失败",
	ERR_KAFKA_PUSH_MSG: "kafka生产消息失败",
	ERR_KAFKA_PULL_MSG: "kafka消费消息失败",
}

func GetMsg(code int) string {
	msg, ok := errMsgMap[code]
	if ok {
		return msg
	}
	return errMsgMap[ERROR]
}
