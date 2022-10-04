package initializers

import "buneydi.com/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
