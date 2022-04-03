package middleware

import (
	"strings"

	"github.com/Eronwin/gin-vue/common"
	"github.com/Eronwin/gin-vue/model"
	"github.com/Eronwin/gin-vue/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get authorization header
		tokenString := ctx.GetHeader("Authorization")

		//validate token format

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Unauthorized(ctx, nil, "权限不足")
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Unauthorized(ctx, nil, "权限不足")
			ctx.Abort()
			return
		}

		//验证通过
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		//用户
		if user.ID == 0 {
			response.Unauthorized(ctx, nil, "权限不足")
			ctx.Abort()
			return
		}

		//用户存在，将user信息放入上下文

		ctx.Set("user", user)
		ctx.Next()
	}
}
