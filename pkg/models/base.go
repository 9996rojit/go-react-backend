package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	Id        string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewV4().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}
