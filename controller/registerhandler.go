package controller

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	// Get the parameter
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	// Data verification

	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "The phone num must be 11 digits",
		})
		return
	}
	// verify the password
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "The password must be at least 6 characters",
		})
		return
	}
	// verify the name
	if len(name) == 0 {
		name = RandomString(10)
		//fmt.Println(name)
	}

	log.Println("name:", name, "phone:", phone, "password:", password)
	// Create a new user
	// Return result
	ctx.JSON(200, gin.H{
		"msg": "Register success!",
	})
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}
