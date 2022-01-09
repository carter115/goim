package router

import (
	"github.com/gin-gonic/gin"
	"gmimo/comet/stub"
	"gmimo/common/e"
	"net/http"
)

// @Summary 首页
// @Tags Home
// @Success 200
// @Router / [get]
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, e.NewResp(e.SUCCESS, nil))
}

// @Summary hello rpc
// @Tags Home
// @param name query string true "名字"
// @Success 200
// @Router /hello [get]
func Hello(c *gin.Context) {
	name := c.Query("name")
	result, errCode := stub.CallHello(c, name)

	c.JSON(http.StatusOK, e.NewResp(errCode, result))
}
