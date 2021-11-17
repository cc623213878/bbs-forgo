package po

import (
	"gorm.io/gorm"
	"time"
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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
