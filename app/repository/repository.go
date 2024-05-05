package repository

import (
	"errors"

	"github.com/nbompetsis/gin-list-notes/app/models"
	"gorm.io/gorm"
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

type ListNotesRepositoryImpl struct {
	DB *gorm.DB
}

func NewListNotesRepositoryImpl(db *gorm.DB) ListNotesRepository {
	return &ListNotesRepositoryImpl{DB: db}
}

func (repo ListNotesRepositoryImpl) Save(list models.List) error {
	result := repo.DB.Create(&list)
	if result.Error != nil {
		return errors.New("list not created")
	}
	return nil
}

func (repo ListNotesRepositoryImpl) Update(list *models.List) error {
	result := repo.DB.Save(&list)
	if result.Error != nil {
		return errors.New("list not updated")
	}
	return nil
}

func (repo ListNotesRepositoryImpl) FindByID(listID uint) (l models.List, err error) {
	var list models.List
	result := repo.DB.Preload("Notes").First(&list, "id = ?", listID)
	if result.Error != nil {
		return list, errors.New("list not found")
	}
	return list, nil
}

func (repo ListNotesRepositoryImpl) FindByOwnerAndActive(owner string, active bool) (lists []models.List, err error) {
	var l []models.List
	result := repo.DB.Preload("Notes").Where("owner = ? and active = ?", owner, active).Find(&l)
	if result.Error != nil || result.RowsAffected == 0 {
		return l, errors.New("lists not found")
	}
	return l, nil
}

func (repo ListNotesRepositoryImpl) Delete(listID uint) error {
	var list models.List
	result := repo.DB.Where("id = ?", listID).Delete(&list)
	if result.Error != nil {
		return errors.New("lists not found")
	}
	return nil
}

func (repo ListNotesRepositoryImpl) Disable(listID uint) error {
	var list models.List
	result := repo.DB.Model(list).Where("id = ? and active = true", listID).Update("active", false)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("list is not active")
	}
	return nil
}

func (repo ListNotesRepositoryImpl) CheckedNote(list *models.List, noteID uint) error {
	result := repo.DB.Model(&models.ListNotes{}).Where("list_id = ? AND note_id = ? AND checked = false", list.ID, noteID).Update("checked", true)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("note is already checked or not found")
	}
	return nil
}

func (repo ListNotesRepositoryImpl) CheckedAllNotes(list *models.List) error {
	var noteIDs []uint
	for _, n := range list.Notes {
		noteIDs = append(noteIDs, n.ID)
	}
	result := repo.DB.Model(&models.ListNotes{}).Where("list_id = ? AND note_id in ? AND checked = false", list.ID, noteIDs).Update("checked", true)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("note is already checked or not found")
	}
	return nil
}
