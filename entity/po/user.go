package po

import (
	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"primaryKey"`
	Password string
	Delete   bool `gorm:"index"`
}

type UserInfo struct {
	Username  string `gorm:"primaryKey"`
	Phone     string
	Email     string
	CreatedAt int64          `gorm:"autoCreateTime"`
	UpdatedAt int64          `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
