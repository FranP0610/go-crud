package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func getDbConnectionChain() string {
	var connectionChain string
	connectionChain = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	return connectionChain
}

func Connection() {
	connectionChain := getDbConnectionChain()
	//log.Println(ConnectionChain)
	var err error
	DB, err = gorm.Open(postgres.Open(connectionChain), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}
}
