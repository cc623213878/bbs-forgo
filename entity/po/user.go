package po

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	Username string `gorm:"size:24;not null;comment:'用户名'"`
	Password string `gorm:"size:24;not null;comment:'密码'"`
	Delete   bool   `gorm:"index;comment:'是否删除【1：已删除】'" json:"delete,omitempty"`
}

type UserInfo struct {
	ID        uint
	UserID    uint           `gorm:"not null;comment:'用户ID'"`
	Phone     string         `gorm:"size:20;comment:'手机号码'"`
	Email     string         `gorm:"size:32;comment:'邮件'"`
	CreatedAt int64          `gorm:"autoCreateTime;index"`
	UpdatedAt int64          `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
