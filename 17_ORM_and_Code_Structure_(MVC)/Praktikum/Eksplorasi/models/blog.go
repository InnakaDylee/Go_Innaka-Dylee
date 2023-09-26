package models

import (
	"github.com/jinzhu/gorm"
)

type Blog struct {
	gorm.Model
	UsersId uint `json:"users_id" form:"users_id"`
	User User `json:"user" gorm:"foreignkey:UsersId"`
	Title string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}