package users

import (
	"bbs-forgo/utils/response"
	"github.com/gin-gonic/gin"
)

// Login 登录
func Login(c *gin.Context) {
	response.Success(c, "")
}

// Logout 退出登录
func Logout(c *gin.Context) {
	response.Success(c, "")
}
