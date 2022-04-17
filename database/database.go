package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBInstance struct {
	DB *gorm.DB
}

var DB DBInstance

func GetConnection() {
	dsn := "host=localhost user=postgres password=150903 dbname=nearfooddb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("Connected database")

	DB = DBInstance{
		DB: db,
	}
}
