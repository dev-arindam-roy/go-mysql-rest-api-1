package models

import "gorm.io/gorm"

type Post struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"` // custom primary key
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	gorm.Model
}

// TableName overrides the table name used by GORM
func (Post) TableName() string {
	return "blog_posts" // custom table name
}
