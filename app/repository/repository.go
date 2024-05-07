package repository

import (
	"github.com/nbompetsis/gin-list-notes/app/models"
)

type ListNotesRepository interface {
	Save(list models.List) error
	Update(listId uint, updatedList models.List) error
	AddListNotes(listID uint, notes []models.Note) error
	FindListNotesByListID(listID uint) (listNotesInfo models.ListNotesInfo, err error)
	FindListNotesByOwner(owner string) (listNotesInfo []models.ListNotesInfo, err error)
	DeleteList(listID uint) error
	CheckListNote(listID uint, noteID uint) error
	CheckListAllNotes(listID uint) error
}
