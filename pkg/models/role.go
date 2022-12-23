package models

import "github.com/9996rojit/backend_go/pkg/config"

// import "gorm.io/gorm"

type Role struct {
	Base
	Name         string `json:"role_name"`
	PermissionId string
	Permission   Permission `gorm:"foreignKey:PermissionId"`
}

func init() {
	config.Connect()
	db = config.GetDb()

}

func GetRole() []Role {
	var role []Role
	db.Model(&Role{}).Find(&role)
	return role
}

func (r *Role) CreateRole() *Role {
	db.Model(&Role{}).Create(&r)
	return r
}
func GetRoleById(id string) *Role {
	var role Role
	db.Model(&Role{}).Where("ID = ?", id).Find(&role)
	return &role
}
func DeleteRoleById(id string) *Role {
	var role *Role
	db.Model(&Role{}).Delete(&role)
	return role
}
