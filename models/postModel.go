package models

import "gorm.io/gorm"

const (
	DRAFT     = 0
	PUBLISHED = 1
)

type Post struct {
	gorm.Model
	Title      string
	Data       string
	Slug       string `gorm:"unique"`
	CoverImage string
	State      uint `gorm:"default:0"`
	UserID     uint
	Comments   []Comment
	Tags       []*Tag `gorm:"many2many:post_tags;"`
	Images     []Image
	PostViews  []PostView
	PostLike   []PostLike
}
