package main

import "time"

type List struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"size:100"`
	Owner   string `gorm:"size:100"`
	Created time.Time
	Active  bool
	Notes   []Note `gorm:"many2many:list_notes;"`
}

type Note struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
}

type ListNotes struct {
	ListID  uint `gorm:"primaryKey"`
	NoteID  uint `gorm:"primaryKey"`
	Checked bool
}
