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
	t := time.Now()
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(t.Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User

		result := initializers.DB.First(&user, claims["sub"])

		if result.Error != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		var session models.Session
		result = initializers.DB.First(&session, "UserID=?", user.ID)
		if result.Error != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		if session.Token != tokenString {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		if session.ExpiresAt.Unix() < t.Unix() {
			initializers.DB.Delete(session, "UserID=?", user.ID)

			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		ctx.Set("user", user)
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.Next()
}
