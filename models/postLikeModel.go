package models

import "gorm.io/gorm"

type PostLike struct {
	gorm.Model
	UserID uint
	PostID uint
}
