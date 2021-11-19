package routers

import (
	"bbs-forgo/middleware"
	"bbs-forgo/services/users"
	"github.com/gin-gonic/gin"
)

func Users(e *gin.Engine) {
	noneLoginGroup := e.Group("/users")
	{
		noneLoginGroup.POST("/login", users.Login)
		noneLoginGroup.POST("/register", users.Register)
	}
	userGroup := e.Group("/users")
	userGroup.Use(middleware.JWTAuth())
	{
		userGroup.POST("/logout", users.Logout)
	}
}
