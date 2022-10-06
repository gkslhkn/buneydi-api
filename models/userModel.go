package models

import "gorm.io/gorm"

const (
	USER   = 0
	AUTHOR = 1
	ADMIN  = 2
)

type User struct {
	gorm.Model
	Name          string
	Email         string `gorm:"unique"`
	UserName      string `gorm:"unique"`
	Password      string
	EmailVerified bool
	Image         string
	Role          uint `gorm:"default:0"`
	UserDetails   UserDetails
	Sessions      []Session
	Posts         []Post
	Images        []Image
	Comments      []Comment
	PostLikes     []PostLike
	CommentLikes  []CommentLike
	PostViews     []PostView
	Tags          []Tag
}

type UserDetails struct {
	gorm.Model
	Address    string
	City       string
	Country    string
	PostalCode string
	Bio        string
	Iban       string
	UserID     uint
}
