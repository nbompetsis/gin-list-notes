package repository

import (
	"errors"
	"fmt"

	"github.com/nbompetsis/gin-list-notes/app/config"
	"github.com/nbompetsis/gin-list-notes/app/models"
	"gorm.io/gorm"
)

type ListNotesRepositoryImpl struct {
	DB *gorm.DB
}

func NewListNotesRepositoryImpl() ListNotesRepository {
	db := config.Connection()
	db.AutoMigrate(&models.List{}, &models.Note{})
	return &ListNotesRepositoryImpl{DB: db}
}

func (repo ListNotesRepositoryImpl) Save(list models.List) error {
	// Check if each note exists
	mergeNotes, err := mergeNotes(repo.DB, list.Notes)
	if err != nil {
		return err
	}
	list.Notes = mergeNotes
	result := repo.DB.Create(&list)
	if result.Error != nil {
		return errors.New("list not created")
	}
	return nil
}

func mergeNotes(DB *gorm.DB, notes []models.Note) ([]models.Note, error) {
	var noteNames []string
	for _, n := range notes {
		noteNames = append(noteNames, n.Name)
	}
	existedNotes, err := findExistedNote(DB, noteNames)
	if err != nil {
		return []models.Note{}, err
	}

	set := make(map[string]bool)
	for _, element := range existedNotes {
		set[element.Name] = true
	}
	var mergeNotes []models.Note
	if len(existedNotes) == len(notes) {
		mergeNotes = existedNotes
	} else {
		for _, element := range noteNames {
			if !set[element] {
				mergeNotes = append(mergeNotes, models.Note{Name: element})
			}
		}
		mergeNotes = append(mergeNotes, existedNotes...)
	}
	return mergeNotes, nil
}

func findExistedNote(DB *gorm.DB, noteNames []string) ([]models.Note, error) {
	var existedNotes []models.Note
	if err := DB.Where("name in ? ", noteNames).Find(&existedNotes).Error; err != nil {
		return []models.Note{}, errors.New("could not retrive notes")
	}
	return existedNotes, nil
}

func (repo ListNotesRepositoryImpl) Update(listID uint, list models.List) error {
	result := repo.DB.Model(&models.List{}).Where("id = ?", listID).Updates(map[string]interface{}{"name": list.Name, "active": list.Active})
	if result.Error != nil {
		return errors.New("list not updated")
	}
	return nil
}

func (repo ListNotesRepositoryImpl) AddListNotes(listID uint, notes []models.Note) error {
	mergeNotes, err := mergeNotes(repo.DB, notes)
	if err != nil {
		return err
	}
	var list models.List
	repo.DB.Where("id=?", listID).Find(&list)

	err = repo.DB.Model(&list).Association("Notes").Replace(mergeNotes)
	if err != nil {
		return err
	}
	return nil
}

func (repo ListNotesRepositoryImpl) FindListNotesByListID(listID uint) (models.ListNotesInfo, error) {
	var listNotes []models.ListNotesInfo
	result := repo.DB.Raw("SELECT lists.id AS list_id, lists.name AS list_name, notes.id AS note_id, "+
		"notes.name AS note_name, list_notes.checked AS note_checked FROM lists "+
		"INNER JOIN list_notes ON lists.id = list_notes.list_id "+
		"INNER JOIN notes ON list_notes.note_id = notes.id WHERE lists.id = ?", listID).Scan(&listNotes)
	if result.Error != nil || result.RowsAffected == 0 || len(listNotes) != 1 {
		return models.ListNotesInfo{}, errors.New("list not found")
	}
	return listNotes[0], nil
}

func (repo ListNotesRepositoryImpl) FindListNotesByOwner(owner string) ([]models.ListNotesInfo, error) {
	var listNotes []models.ListNotesInfo
	result := repo.DB.Raw("SELECT lists.id AS list_id, lists.name AS list_name, notes.id AS note_id, "+
		"notes.name AS note_name, list_notes.checked AS note_checked FROM lists "+
		"INNER JOIN list_notes ON lists.id = list_notes.list_id "+
		"INNER JOIN notes ON list_notes.note_id = notes.id WHERE lists.owner = ?", owner).Scan(&listNotes)
	if result.Error != nil || result.RowsAffected == 0 {
		return []models.ListNotesInfo{}, errors.New("lists not found")
	}
	return listNotes, nil
}

func (repo ListNotesRepositoryImpl) DeleteList(listID uint) error {
	var list models.List
	result := repo.DB.Where("id = ?", listID).Delete(&list)
	if result.Error != nil {
		return errors.New("list not found")
	}
	return nil
}

func (repo ListNotesRepositoryImpl) CheckListNote(listID uint, noteID uint) error {
	result := repo.DB.Model(&models.ListNotes{}).Where("list_id = ? AND note_id = ? AND checked = false", listID, noteID).Update("checked", true)
	if result.Error != nil || result.RowsAffected == 0 {
		return fmt.Errorf("note %d is already checked or not found for the list %d", noteID, listID)
	}
	return nil
}

func (repo ListNotesRepositoryImpl) CheckListAllNotes(listID uint) error {
	result := repo.DB.Model(&models.ListNotes{}).Where("list_id = ?", listID).Update("checked", true)
	if result.Error != nil || result.RowsAffected == 0 {
		return fmt.Errorf("notes are already checked or list %d not found", listID)
	}
	return nil
}
