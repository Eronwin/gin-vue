package main

import (
	"net/http"

	"github.com/Eronwin/gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.GET("/ping", controller.Ping)
	r.Run(":9090") // 这边我们使用9999 这个端口
}
