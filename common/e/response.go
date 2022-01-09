package e

type Response struct {
	code   int
	msg    string
	result interface{}
}

// 统一响应格式
func NewResp(code int, result interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":   code,
		"msg":    GetMsg(code),
		"result": result,
	}
}
