package article

import (
	"bbs-forgo/db"
	"bbs-forgo/entity/dto"
	"bbs-forgo/entity/po"
	"bbs-forgo/entity/vo"
	"bbs-forgo/log"
	"bbs-forgo/utils/response"
	"github.com/gin-gonic/gin"
)

func GetArticleList(c *gin.Context) {
	var page dto.Page
	err := c.BindJSON(&page)
	if err != nil {
		log.GetLogger().Error(err.Error())
		response.Error(c, err.Error())
		return
	}
	var result []vo.ArticleList
	var count int64
	db.DB.Model(&po.Article{}).
		Limit(page.PageSize).
		Offset((page.PageNum - 1) * page.PageSize).
		Order("updated_at desc").Find(&result)
	db.DB.Model(&po.Article{}).Count(&count)
	response.Success(c, response.ResponsePage(page.PageNum, page.PageSize, count, &result))
}
