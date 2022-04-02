package controller

import (
	"log"
	"net/http"

	"github.com/Eronwin/gin-vue/common"
	"github.com/Eronwin/gin-vue/model"
	"github.com/Eronwin/gin-vue/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//用户注册
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
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "加密失败"})
		return
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)

	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

//用户登录
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
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
	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该用户不存在"})
		return
	}
	//验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("[token generate error : %v]", err)
		return
	}

	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
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
