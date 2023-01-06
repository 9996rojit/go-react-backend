package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	db_string := "host=localhost user=admin password=admin dbname=oms_backend port=5432 sslmode=disable TimeZone=Asia/Kathmandu"
	d, err := gorm.Open(postgres.Open(db_string), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
