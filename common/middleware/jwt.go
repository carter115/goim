package middleware

import (
	"github.com/gin-gonic/gin"
	"gmimo/comet/util"
	"gmimo/common/e"
	"gmimo/common/log"
	"net/http"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var result interface{}

		code = e.SUCCESS
		token := c.Request.Header.Get("token") //header
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			// 解析token，检查是否有效
			_, err := util.ParseToken(token)
			if err != nil {
				code = e.ERR_AUTH_TOKEN_FAIL
				result = err.Error()
			}
		}

		if code != e.SUCCESS {
			log.Warnc(c, e.GetMsg(code))
			c.JSON(http.StatusUnauthorized, e.NewResp(code, result))
			c.Abort()
			return
		}
		c.Next()
	}
}
