package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey,autoIncrement"`
	Name  string `form:"name" json:"name" validate:"required"`
	Email string `form:"email" json:"email" validate:"required,email"`
}

func ListAllUsers(db *gorm.DB) []User {
	var users []User
	db.Order("id desc").Find(&users)
	return users
}
