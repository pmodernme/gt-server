package model

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	UserName    string `gorm:"type:varchar(100);unique_index" json:"username"`
	DisplayName string `json:"display_name"`
}
