package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID       int    `json:"id,omitempty"`
	Name     string `gorm:"type:varchar(255);not null",json:"name,omitempty"`
	Password string `gorm:"type:varchar(100);not null",jjson:"password,omitempty"`
	Phone    string `gorm:"size:255;not null",jjson:"phone,omitempty"`
}
