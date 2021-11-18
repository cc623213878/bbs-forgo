package users

import (
	"bbs-forgo/db"
	"bbs-forgo/entity/dto"
	"bbs-forgo/entity/po"
	"bbs-forgo/log"
	"bbs-forgo/utils/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login 登录
func Login(c *gin.Context) {
	var user dto.User
	err := c.BindJSON(&user)
	if err != nil {
		log.GetLogger().Error("param error")
		return
	}
	if db.DB.First(&user).RowsAffected > 0 {

	}
	response.Success(c, "")
}

// Logout 退出登录
func Logout(c *gin.Context) {
	response.Success(c, "")
}

func Register(c *gin.Context) {
	var user dto.UserInfo
	err := c.BindJSON(&user)
	if err != nil {
		log.GetLogger().Error(err.Error())
		response.Error(c, err.Error())
		return
	}
	var userPo po.User
	result := db.DB.First(&userPo, "username = ?", user.Username)
	if result.RowsAffected >= 0 {
		log.GetSugarLogger().Warn("Register[", user.Username, "]该用户名已存在")
		response.Error(c, "该用户名已存在！")
		return
	}
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&po.User{
			Username: user.Username,
			Password: user.Password,
		}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(&po.UserInfo{
			Username: user.Username,
			Phone:    user.Phone,
			Email:    user.Email,
		}).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		log.GetSugarLogger().Error(err.Error(), "注册失败！")
		response.Error(c, "注册失败！")
		return
	}
	response.Success(c, gin.H{"username": user.Username})
}
