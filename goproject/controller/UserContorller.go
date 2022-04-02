package controller

import (
	"net/http"

	"github.com/Eronwin/gin-vue/common"
	"github.com/Eronwin/gin-vue/model"
	"github.com/Eronwin/gin-vue/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	//数据验证
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不小于6位"})
		return
	}
	if len(name) == 0 {
		name = util.RandomStr(10)
	}

	if isPhoneExist(DB, phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该用户已经存在"})
		return
	}

	//创建用户

	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	DB.Create(&newUser)

	//返回结果
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
