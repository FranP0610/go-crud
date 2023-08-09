package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dsn = "host=localhost user=admin password=password dbname=gorm port=5432"
var DB *gorm.DB

func Connection() {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}
}
