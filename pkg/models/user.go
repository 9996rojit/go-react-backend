package models

import (
	"fmt"

	"github.com/9996rojit/backend_go/pkg/config"
	"github.com/9996rojit/backend_go/pkg/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	Base
	RoleId    string `json:"role_id"`
	Password  string `json:"password"`
	Name      string `json:"user_name"`
	Email     string `gorm:"unique" json:"user_email"`
	Contact   string `json:"user_contact"`
	Location  string `json:"user_location"`
	Gender    string `json:"user_gender"`
	Age       int    `json:"user_age"`
	CompanyId string `json:"user_company_id"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&User{}, &Company{}, &Product{}, &ProductImage{}, &Role{})
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashPassword, err := utils.HashedPassword(u.Password)
	if err != nil {
		fmt.Print("There is error while hashing password")
	}
	u.Password = hashPassword
	return
}

func GetUser() []User {
	var user []User
	db.Model(&User{}).Find(&user)
	return user
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

//find user by email

func FindUserByEmail(email string) *User {
	var user User
	db.Model(&User{}).Where("email =?", email).Find(&user)
	return &user
}
