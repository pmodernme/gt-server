package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(100);unique_index"`
}
