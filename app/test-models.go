package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=password dbname=notes port=5432 sslmode=disable TimeZone=Europe/Athens"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&List{}, &Note{}, &ListNotes{})
	note := Note{Name: "Example"}
	db.Create(&note)
}