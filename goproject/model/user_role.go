package model

import "github.com/jinzhu/gorm"

type UserRole struct {
	gorm.Model
	Userid string `gorm:"varchar(255);not null;"`
	Roleid string `gorm:"varchar(255);not null;"`
}
