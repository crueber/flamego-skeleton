package models

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  ID uint `gorm:"primaryKey,autoIncrement"`
  Name string `form:"name" json:"name" validate:"required"`
  Email string `form:"email" json:"email" validate:"required,email"`
}

func (u *User) Save(db *gorm.DB) error {
  result := db.Save(&u)
  return result.Error
}

func (u *User) Delete(db *gorm.DB) error {
  result := db.Delete(&u)
  return result.Error
}

func ListAllUsers(db *gorm.DB) []User {
  var users []User
  db.Find(&users)
  return users
}

func GetUser(db *gorm.DB, id int) User {
  var user User
  db.First(&user, id)
  return user
}

