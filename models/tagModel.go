package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Data   string `gorm:"unique"`
	Slug   string `gorm:"unique"`
	UserID uint
	Posts  []*Post `gorm:"many2many:post_tags;"`
}
