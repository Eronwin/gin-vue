package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}
func Fail(ctx *gin.Context, code int, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
func UnprocessableEntity(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusUnprocessableEntity, 422, data, msg)
}
func InternalServerError(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusInternalServerError, 500, data, msg)
}
func BadRequest(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusBadRequest, 400, data, msg)
}
func Unauthorized(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusUnauthorized, 401, data, msg)
}
