package models

import "gorm.io/gorm"

type VerificationCode struct {
	gorm.Model
	Email            string
	VerificationCode uint
}
