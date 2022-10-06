package models

import "gorm.io/gorm"

type CommentLike struct {
	gorm.Model
	CommentID uint
	UserID    uint
}
