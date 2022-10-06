package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Data   string
	UserID uint
	PostID uint
}
