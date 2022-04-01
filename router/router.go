package router

import (
	"net/http"

	"github.com/Eronwin/gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	// r.GET("/ping", controller.Ping)
	r.POST("/api/user/register", controller.Register)
	r.Run(":9090")
}
