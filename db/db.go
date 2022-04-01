package db

import (
	"fmt"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	database := "gin_vue"
	username := "root"
	password := "canku555"
	charset := "utf8"
	port := "3306"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		charset,
		database)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}
	return db
}

func Run() {
	DB := db.InitDB()
	defer DB.Close()
}
