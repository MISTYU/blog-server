package model

import (
	"time"

	"gorm.io/gorm"
)

// 基础结构
type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

// index:idx_article 自动生成索i引
// ArticleId   string `gorm:"index:idx_article;unique;type:varchar(11);not null"`
// type Blog struct {
// 	BaseModel
// 	Title       string `gorm:"type:varchar(100);not null"`
// 	ArticleId   string `gorm:"index:idx_article;unique;type:varchar(11);not null"`
// 	Tag         string `gorm:"type:varchar(24);not null"`
// 	Description string `gorm:"type:varchar(255);not null"`
// 	Content     string `gorm:"type:text;not null"`
// }

type Blog struct {
	BaseModel
	Title        string `gorm:"type:varchar(100);not null"`
	Tag          string `gorm:"type:varchar(24);not null"`
	Description  string `gorm:"type:varchar(255);not null"`
	ArticleId    string `gorm:"index:idx_article;unique;type:varchar(11);not null"`
	BlogdetailID int
	Blogdetail   Blogdetail
}

type Blogdetail struct {
	ID      int
	Content string `gorm:"type:text;not null"`
}
