package controller

import (
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "register",
	})
}
