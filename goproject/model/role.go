package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Rolename string `gorm:"varchar(255);not null;unique"`
}
