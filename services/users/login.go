package users

import (
	"bbs-forgo/entity/dto"
	"bbs-forgo/log"
	"bbs-forgo/utils/response"
	"github.com/gin-gonic/gin"
)

// Login 登录
func Login(c *gin.Context) {
	var user dto.User
	err := c.BindJSON(&user)
	if err != nil {
		log.GetLogger().Error("param error")
		return
	}

	response.Success(c, user)
}

// Logout 退出登录
func Logout(c *gin.Context) {
	response.Success(c, "")
}
