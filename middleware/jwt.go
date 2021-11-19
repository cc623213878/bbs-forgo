package middleware

import (
	"bbs-forgo/log"
	"bbs-forgo/utils/jwt"
	"bbs-forgo/utils/response"
	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Token")
		if token == "" {
			response.Error(c, "无权限访问！")
			c.Abort()
			return
		}
		log.GetSugarLogger().Info("token: ", token)
		j := jwt.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				response.Error(c, "授权已过期！")
				c.Abort()
				return
			}
			log.GetLogger().Error(err.Error())
			response.Error(c, "无权限访问！")
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}
