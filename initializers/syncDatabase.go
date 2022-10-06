package initializers

import "buneydi.com/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.UserDetails{}, &models.Session{}, &models.Post{}, &models.Tag{}, &models.Comment{}, &models.Image{}, &models.VerificationCode{}, &models.PostView{}, &models.PostLike{}, &models.CommentLike{})
}
