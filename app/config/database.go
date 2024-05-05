package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	host := os.Getenv("host")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbName := os.Getenv("dbName")
	port := os.Getenv("port")
	timeZone := os.Getenv("timeZone")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, dbName, port, timeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error creating db connection", err)
	}

	return db
}
