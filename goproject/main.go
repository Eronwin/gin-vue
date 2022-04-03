package main

import (
	"log"
	"os"

	"github.com/Eronwin/gin-vue/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	wordDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(wordDir + "/config")
	log.Println("config path:", wordDir+"/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
