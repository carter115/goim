package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gmimo/common/constant"
	"gmimo/common/log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UnixNano()
		c.Next()

		statusCode := c.Writer.Status()
		delay := (time.Now().UnixNano() - start) / int64(time.Millisecond) //单位:毫秒
		clientIP := c.ClientIP()
		method := c.Request.Method
		uri := c.Request.RequestURI

		// 从gin.Content拿取body
		var body []byte
		cbody := c.Value(gin.BodyBytesKey) // 经过校验的body体
		if cbody != nil {
			body = cbody.([]byte)
			if len(body) > 10*constant.KB { // body最多打印10KB的日志
				body = body[:10*constant.KB]
			}
		}

		userAgent := c.Request.UserAgent()

		log.Infoc(c, logMessage{statusCode, delay, clientIP, method,
			uri, string(body), userAgent})
	}
}

// 一条日志消息
type logMessage struct {
	statusCode int
	delay      int64
	client     string
	method     string
	uri        string
	body       string
	userAgent  string
}

// ELK 格式化输出
func (m logMessage) String() string {
	return fmt.Sprintf("Status:%d Time:%dms Client:%s\n"+
		"Method:%s\n"+
		"Uri:%s\n"+
		"Body:%s\n"+
		"UserAgent:%s",
		m.statusCode, m.delay, m.client, m.method, m.uri,
		m.body, m.userAgent,
	)
}
