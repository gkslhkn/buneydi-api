package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Token     string `gorm:"unique"`
	ExpiresAt time.Time
	UserID    uint
}
