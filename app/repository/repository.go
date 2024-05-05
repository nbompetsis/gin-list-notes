package repository

import (
	"errors"

	"github.com/nbompetsis/gin-list-notes/app/models"
	"gorm.io/gorm"
)

type ListNotesRepository interface {
	Save(list models.List) error
	Update(list *models.List) error
	FindById(listId uint) (list models.List, err error)
	FindByOwnerAndActive(owner string, active bool) (lists []models.List, err error)
	Delete(listId uint) error
	Disable(listId uint) error
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

func (repo ListNotesRepositoryImpl) FindById(listId uint) (l models.List, err error) {
	var list models.List
	result := repo.DB.First(&list, "id = ?", listId)
	if result.Error != nil {
		return list, errors.New("list not found")
	}
	return list, nil
}

func (repo ListNotesRepositoryImpl) FindByOwnerAndActive(owner string, active bool) (lists []models.List, err error) {
	var l []models.List
	result := repo.DB.Where("owner = ? and active = ?", owner, active).Find(&l)
	if result.Error != nil || result.RowsAffected == 0 {
		return l, errors.New("lists not found")
	}
	return l, nil
}

func (repo ListNotesRepositoryImpl) Delete(listId uint) error {
	var list models.List
	result := repo.DB.Where("id = ?", listId).Delete(&list)
	if result.Error != nil {
		return errors.New("lists not found")
	}
	return nil
}

func (repo ListNotesRepositoryImpl) Disable(listId uint) error {
	var list models.List
	result := repo.DB.Model(list).Where("id = ? and active = true", listId).Update("active", false)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("list is not active")
	}
	return nil
}
