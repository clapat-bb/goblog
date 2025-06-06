package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:text;not null" json:"context"`
	Content     string         `gorm:"type:text;not null" json:"content"`
	UserID      uint           `json:"user_id"`
	User        User           `json:"author"`
	IsDraft     bool           `gorm:"default:false" json:"is_draft"`
	IsTop       bool           `gorm:"default:false" json:"is_top"`
	IsRecommend bool           `gorm:"default:false" json:"is_recommend"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Tags        []*Tag         `gorm:"many2many:post_tags;" json:"tags"` // 该标签通过 gorm:"many2many:post_tags;" 指定了用 post_tags 中间表来建立 Post 与 Tag 的多对多关联，并在 JSON 序列化时将该字段命名为 tags。
}
