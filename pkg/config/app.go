package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	db_string := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Kathmandu"
	d, err := gorm.Open(postgres.Open(db_string), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
	fmt.Printf("this is database connection test in go lang is \n %w \n", d)
}

func GetDb() *gorm.DB {
	return db
}
