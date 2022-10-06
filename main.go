package main

import (
	"fmt"

	"buneydi.com/api/controllers"
	"buneydi.com/api/initializers"
	"buneydi.com/api/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	fmt.Println("Starting Server...")
	r := gin.Default()
	r.POST("/users/user/signup", controllers.SignUp)
	r.POST("/users/user/login", controllers.LogIn)
	r.POST("/users/user/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/users/user/author/:id", controllers.GetAuthorInformation)
	r.Run()
}
