package middleware

import (
	"net/http"
	"os"
	"time"

	"buneydi.com/api/initializers"
	"buneydi.com/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(ctx *gin.Context) {

	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		ctx.Set("user", user)
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.Next()
}
