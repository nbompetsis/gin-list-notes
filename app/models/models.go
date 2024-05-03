package models

import "time"

type List struct {
	ID      uint      `gorm:"primaryKey"`
	Name    string    `gorm:"size:100;not null"`
	Owner   string    `gorm:"size:100;not null"`
	Created time.Time `gorm:"default:CURRENT_TIMESTAMP;not null"`
	Active  bool      `gorm:"default:true;not null"`
	Notes   []Note    `gorm:"many2many:list_notes;"`
}

type Note struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;unique;not null"`
}

type ListNotes struct {
	Checked bool `gorm:"default:false;not null"`
}
