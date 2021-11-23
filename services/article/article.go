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
	var result []vo.Article
	var count int64
	xt := db.DB.Model(&po.Article{}).
		Limit(page.PageSize).
		Offset((page.PageNum - 1) * page.PageSize).
		Order("updated_at desc").Find(&result)
	if xt.Error != nil {
		response.Error(c, "系统错误！")
	}
	xt = db.DB.Model(&po.Article{}).Count(&count)
	if xt.Error != nil {
		response.Error(c, "系统错误！")
	}
	response.Success(c, response.ResponsePage(page.PageNum, page.PageSize, count, &result))
}

func GetArticleByID(c *gin.Context) {
	var article dto.Article
	err := c.BindJSON(&article)
	if err != nil {
		log.GetLogger().Error(err.Error())
		response.Error(c, err.Error())
		return
	}
	var result vo.Article
	xt := db.DB.Model(&po.Article{}).Where(&article).Find(&result)
	if xt.Error != nil {
		response.Error(c, "系统错误！")
	}
	response.Success(c, result)
}

func WriteArticle(c *gin.Context) {

}
