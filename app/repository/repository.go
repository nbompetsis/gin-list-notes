package repository

import (
	"github.com/nbompetsis/gin-list-notes/app/models"
)

type ListNotesRepository interface {
	Save(list models.List) error
	Update(list *models.List) error
	FindByID(listID uint) (list models.List, err error)
	FindByOwnerAndActive(owner string, active bool) (lists []models.List, err error)
	Delete(listID uint) error
	Disable(listID uint) error
	CheckedNote(list *models.List, noteID uint) error
	CheckedAllNotes(list *models.List) error
}
