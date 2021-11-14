package routers

import (
	"bbs-forgo/services/users"
	"github.com/gin-gonic/gin"
)

func Login(e *gin.Engine) {
	usersGroup := e.Group("/users")
	{
		usersGroup.POST("/login", users.Login)
	}
}
