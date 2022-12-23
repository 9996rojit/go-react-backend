package models

import "gorm.io/gorm"

type Permission struct {
	Base
	Name string `json:"permission_name"`
}

func GetPermissions() []Permission {
	var AllPermissions []Permission
	db.Model(&Permission{}).Find(&AllPermissions)
	return AllPermissions
}

func (p *Permission) CreatePermisssion() *Permission {
	db.Create(&p)
	return p
}

func GetPermissionById(id string) (*Permission, *gorm.DB) {
	var permission Permission
	db.Model(&Permission{}).Where("Id = ?", id).Find(&permission)
	return &permission, db
}

func DeletePermissionById(id string) *Permission {
	var permission Permission
	db.Model(&Permission{}).Where("Id = ?", id).Delete(&permission)
	return &permission
}
