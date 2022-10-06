package controllers

import (
	"net/http"
	"os"
	"time"

	"buneydi.com/api/initializers"
	"buneydi.com/api/models"
	"buneydi.com/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	//get email and password from req.body
	var body struct {
		Email    string
		Password string
		UserName string
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
			"error": "Failed to read email.",
		})
		return
	}
	if body.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read password.",
		})
		return
	}

	if !utils.CheckEmail(body.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is not valid.",
		})
		return
	}

	//hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}
	//create the user

	var user models.User

	result := initializers.DB.First(&user, "email=?", body.Email)

	if result.Error == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Signed up before.",
		})
		return
	}

	user = models.User{Email: body.Email, Password: string(hash), UserName: body.UserName}
	result = initializers.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
		return
	}
	//respond with 200
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func LogIn(ctx *gin.Context) {
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
			"error": "Failed to read email.",
		})
		return
	}
	if body.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read password.",
		})
		return
	}

	if !utils.CheckEmail(body.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is not valid.",
		})
		return
	}

	var user models.User

	initializers.DB.First(&user, "email=?", body.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password.",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password.",
		})
		return
	}

	exp := time.Now().Add(time.Hour * 24 * 30 * 12)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": exp.Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	var session = &models.Session{
		Token:     tokenString,
		ExpiresAt: exp,
		UserID:    user.ID,
	}

	result := initializers.DB.Where("user_id=?", user.ID).Delete(&models.Session{})
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	result = initializers.DB.Create(&session)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Token not created.",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Token not created.",
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)

	ctx.SetCookie("Authorization", tokenString, 3600*24*30*12, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
