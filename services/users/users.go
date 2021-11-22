package users

import (
	"bbs-forgo/db"
	"bbs-forgo/entity/dto"
	"bbs-forgo/entity/po"
	"bbs-forgo/log"
	myjwt "bbs-forgo/utils/jwt"
	"bbs-forgo/utils/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// Login 登录
func Login(c *gin.Context) {
	var user dto.User
	err := c.BindJSON(&user)
	if err != nil {
		log.GetLogger().Error("param error")
		return
	}
	res := db.DB.First(&user)
	if res != nil && res.RowsAffected > 0 {
		expiresTime := time.Now().Unix() + 3600
		j := myjwt.NewJWT()
		claims := myjwt.CustomClaims{
			Username: user.Username,
			Password: user.Password,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expiresTime,
				// 指定token发行人
				Issuer: "go-forbbs",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			log.GetLogger().Error("CreateToken error")
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		response.Success(c, gin.H{"token": token})
		return
	}
	response.Error(c, "用户名或密码错误！")
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
		insertUser := po.User{
			Username: user.Username,
			Password: user.Password,
		}
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&insertUser).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(&po.UserInfo{
			UserID: insertUser.ID,
			Phone:  user.Phone,
			Email:  user.Email,
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
