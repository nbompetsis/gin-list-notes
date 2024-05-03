package main

import (
	"log"

	"github.com/nbompetsis/gin-list-notes/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=password dbname=notes port=5432 sslmode=disable TimeZone=Europe/Athens"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.List{}, &models.Note{}, &models.ListNotes{})

	// notes := []models.Note{
	// 	{Name: "Note1"},
	// 	{Name: "Note2"},
	// }
	// db.Create(&notes)

	// var notes []models.Note
	// db.Where("name in (?,?)", "Note1", "Note2").Find(&notes)
	// for _, note := range notes {
	// 	fmt.Println(note)
	// }

	// lists := []models.List{
	// 	{Name: "listA", Owner: "Nik"},
	// 	{Name: "listB", Owner: "John"},
	// }
	// db.Create(&lists)

	// var notes []models.Note
	// db.Where("name in (?,?)", "Note1", "Note2").Find(&notes)

	// var lists []models.List
	// db.Where("owner = ?", "Nik").Find(&lists)

	// for _, list := range lists {
	// 	list.Notes = notes
	// 	db.Save(list)
	// }
}
