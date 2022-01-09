package router

//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/gin-gonic/gin/binding"
//	"gmimo/common/e"
//	"gmimo/common/log"
//	"net/http"
//)
//
//// @Summary 用户信息
//// @Tags 用户管理
//// @Param username query string true "用户名"
//// @Param token header string true "Token"
//// @Success 200
//// @Router /user [get]
//func User(c *gin.Context) {
//
//	username := c.Request.FormValue("username")
//
//	// TODO 请求参数检验不过，马上返回
//
//	log.Infoc(c, fmt.Sprintf("welcome [%s]", username))
//	c.JSON(200, e.NewResp(e.SUCCESS, nil))
//}
//
//// @Summary 修改用户
//// @Tags 用户管理
//// @Param username query string true "用户名"
//// @Param token header string true "Token"
//// @Success 200
//// @Router /user/update [post]
//func UserUpdate(c *gin.Context) {
//
//	// 1. 支持请求body的Content-Type: form-data,x-www-form-urlencoded
//	//var formData UserUpdateForm
//	//if err := c.ShouldBindQuery(&formData); err != nil {
//	//	log.Warnc(c, err.Error())
//	//	c.JSON(http.StatusBadRequest, e.NewResp(e.INVALID_PARAMS, nil))
//	//	return
//	//}
//
//	// 2. 支持请求body的Content-Type: raw(text,json)
//	var jsonData = UserUpdateForm{}
//	// ShouldBindBodyWith 读取 c.Request.Body 并将结果存入上下文。
//	if err := c.ShouldBindBodyWith(&jsonData, binding.JSON); err != nil {
//		log.Warnc(c, err.Error())
//		c.JSON(http.StatusBadRequest, e.NewResp(e.INVALID_PARAMS, nil))
//		return
//	}
//
//	log.Infoc(c, fmt.Sprintln("update userinfo "))
//	c.JSON(200, e.NewResp(e.SUCCESS, nil))
//}

// ----------------表单Form----------------------

//type UserUpdateForm struct {
//	Username string `form:"username" json:"username" binding:"required"`
//	Password string `form:"password" json:"password" binding:"required"`
//}

//type UserUpdateForm struct {
//	Username string `form:"username" binding:"required"`
//	Password string `form:"password" binding:"required"`
//}
