package po

import "gorm.io/gorm"

type Article struct {
	ID        uint
	Title     string         `gorm:"size:256;not null;comment:'文章名'"`
	Keywords  string         `gorm:"size:256;comment:'关键词'"`
	CoverPath string         `gorm:"size:512;comment:'文章封面'"`
	Content   string         `gorm:"size:1048576;not null;comment:'文章内容'"`
	Username  string         `gorm:"size:24;not null;comment:'用户名'"`
	CreatedAt int64          `gorm:"autoCreateTime;index"`
	UpdatedAt int64          `gorm:"autoUpdateTime;index"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
