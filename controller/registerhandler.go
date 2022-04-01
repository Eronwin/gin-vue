package controller

import (
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	// Get the parameter
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	// Data verification
	// Create a new user
	// Return result
	ctx.JSON(200, gin.H{
		"msg": "Register success!",
	})
}
