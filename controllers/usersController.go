package controllers

import (
	"net/http"

	"buneydi.com/api/initializers"
	"buneydi.com/api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	//get email and password from req.body
	var body struct {
		Email    string
		Password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})

		return
	}
	//check email and password
	if body.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Faied to read email.",
		})
		return
	}
	if body.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Faied to read password.",
		})
		return
	}

	//hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Faied to hash password.",
		})
		return
	}
	//create the user

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Faied to create user.",
		})
		return
	}
	//respond with 200
	ctx.JSON(http.StatusOK, gin.H{})
}

func LogIn(ctx *gin.Context) {

}
