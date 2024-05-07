package repository

import (
	"github.com/nbompetsis/gin-list-notes/app/models"
)

type ListNotesRepository interface {
	Save(list models.List) error
	Update(listId uint, updatedList models.List) error
	FindListNotesByListID(listID uint) (listNotesInfo models.ListNotesInfo, err error)
	FindListNotesByOwner(owner string) (listNotesInfo []models.ListNotesInfo, err error)
	DeleteList(listID uint) error
	CheckNote(listID uint, noteID uint) error
	CheckAllNotes(listID uint) error
}
