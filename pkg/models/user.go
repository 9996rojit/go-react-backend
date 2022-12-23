package models

import (
	"github.com/9996rojit/backend_go/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	Base
	RoleId   string
	Role     Role   `gorm:"foreignKey:RoleId"`
	Name     string `json:"user_name"`
	Email    string `json:"user_email"`
	Contact  string `json:"user_contact"`
	Location string `json:"user_location"`
	Gender   string `json:"user_gender"`
	Age      string `json:"user_age"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&User{})
}

func GetUser() *User {
	var user User
	db.Model(&User{}).Find(&user)
	return &user
}

func (u *User) CreateUser() *User {
	db.Model(&User{}).Create(u)
	return u
}

func GetUserById(id string) (*User, *gorm.DB) {
	var user User
	db.Model(&User{}).Where("ID = ?", id).Find(&user)
	return &user, db
}

func DeleteUser(id string) *User {
	var user User
	db.Model(&User{}).Where("ID = ?", id).Delete(&user)
	return &user

}
