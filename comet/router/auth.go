package router

import (
	"github.com/gin-gonic/gin"
	"gmimo/comet/util"
	"gmimo/common/e"
)

type auth struct {
	Username string
	Password string
}

// @Summary 生成Token
// @Tags 认证
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	var (
		code   = e.SUCCESS
		result string
	)
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	a := auth{username, password}
	if a.Check() {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = e.ERR_AUTH_TOKEN
		} else {
			result = token //通过检验生成token
		}
	} else {
		code = e.ERR_USER_PWD_FAIL
	}

	c.JSON(200, e.NewResp(code, result))
}

func (a auth) Check() bool {
	// TODO 检验帐号和密码
	return true
}
