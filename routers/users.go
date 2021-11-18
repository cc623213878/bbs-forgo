package routers

import (
	"bbs-forgo/services/users"
	"github.com/gin-gonic/gin"
)

func Users(e *gin.Engine) {
	usersGroup := e.Group("/users")
	{
		usersGroup.POST("/login", users.Login)
		usersGroup.POST("/register", users.Register)
	}
}
