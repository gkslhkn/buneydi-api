package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": ctx.Keys["user"],
	})
}
