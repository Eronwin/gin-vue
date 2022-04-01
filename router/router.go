package router

import (
	"github.com/Eronwin/gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.GET("/api/user/register", controller.Register)
}
