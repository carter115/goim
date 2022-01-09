package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gmimo/common/e"
	"gmimo/common/log"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 打印错误日志
				log.Errorc(c, fmt.Sprintf("PANIC: %v", err))
				// panic 错误，统一包装成json数据，进行返回
				c.JSON(http.StatusInternalServerError, e.NewResp(e.ERROR, nil))
				c.Abort()
			}
		}()
		c.Next()
	}
}
