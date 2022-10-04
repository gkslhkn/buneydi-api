package main

import (
	"fmt"

	"buneydi.com/api/controllers"
	"buneydi.com/api/initializers"
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
	r.Run()
}
