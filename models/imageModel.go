package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Url    string `gorm:"unique"`
	PostID uint
	UserID uint
}
