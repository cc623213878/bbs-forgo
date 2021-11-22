package routers

import (
	"bbs-forgo/services/article"
	"github.com/gin-gonic/gin"
)

func Article(e *gin.Engine) {
	articleGroup := e.Group("/article")
	{
		articleGroup.POST("/getArticleList", article.GetArticleList)
	}
}
