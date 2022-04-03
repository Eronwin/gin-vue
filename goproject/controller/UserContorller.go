package controller

import (
	"github.com/Eronwin/gin-vue/common"
	"github.com/Eronwin/gin-vue/dto"
	"github.com/Eronwin/gin-vue/model"
	"github.com/Eronwin/gin-vue/response"
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
		response.UnprocessableEntity(ctx, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.UnprocessableEntity(ctx, nil, "密码不小于6位")
		return
	}
	if len(name) == 0 {
		name = util.RandomStr(10)
	}

	if isPhoneExist(DB, phone) {
		response.UnprocessableEntity(ctx, nil, "该用户已经存在")
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.InternalServerError(ctx, nil, "加密失败")
		return
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)

	//返回结果
	response.Success(ctx, nil, "注册成功")
}

//用户登录
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	//数据验证
	if len(phone) != 11 {
		response.UnprocessableEntity(ctx, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.UnprocessableEntity(ctx, nil, "密码不小于6位")
		return
	}
	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		response.UnprocessableEntity(ctx, nil, "该用户不存在")
		return
	}
	//验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.BadRequest(ctx, nil, "密码错误")
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.InternalServerError(ctx, nil, "系统异常")
		// log.Printf("[token generate error : %v]", err)
		return
	}
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDTO(user.(model.User))}, "获取用户信息成功")
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	return user.ID != 0
}
