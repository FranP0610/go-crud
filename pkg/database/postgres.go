package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func getDbConnectionChain() string {
	var connectionChain string
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("error loading .env file")
	} else {
		connectionChain = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"))
	}
	//connectionChain = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_NAME"),
	//	os.Getenv("DB_PORT"))
	//log.Printf("db user is %s", os.Getenv("DB_USER"))
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
