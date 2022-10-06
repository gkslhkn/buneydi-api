package models

import "gorm.io/gorm"

type PostView struct {
	gorm.Model
	Ip     string
	PostID uint
	UserID uint
}
