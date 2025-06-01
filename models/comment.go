package models

import "gorm.io/gorm"

type Comment struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"` // custom primary key
	Comment string `gorm:"type:text;not null" json:"comment"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	BlogID  uint   `gorm:"not null" json:"blog_id"`
	gorm.Model
}

// TableName overrides the table name used by GORM
func (Comment) TableName() string {
	return "comments" // custom table name
}
