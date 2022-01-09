package e

import "errors"

// 自定义错误类型
func NewError(code int) error {
	return errors.New(GetMsg(code))
}
