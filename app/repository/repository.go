package repository

import (
	"log"

	"github.com/nbompetsis/gin-list-notes/app/models"
	"gorm.io/gorm"
)

type ListNotesRepository interface {
	Save(list models.List)
	FindById(listId uint) (tags models.List, err error)
	FindByOwnerAndActive(owner string, active bool) (lists []models.List, err error)
	Delete(listId uint)
	Disable(listId uint)
}

type ListNotesRepositoryImpl struct {
	DB *gorm.DB
}

func NewListNotesRepositoryImpl(db *gorm.DB) ListNotesRepository {
	return &ListNotesRepositoryImpl{DB: db}
}

func (repo ListNotesRepositoryImpl) Save(list models.List) {
	result := repo.DB.Create(&list)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func (repo ListNotesRepositoryImpl) FindById(listId uint) (l models.List, err error) {
	var list models.List
	result := repo.DB.First(&list, "id = ?", listId)
	if result.Error != nil {
		return list, nil
	} else {
		return list, result.Error
	}
}

func (repo ListNotesRepositoryImpl) FindByOwnerAndActive(owner string, active bool) (lists []models.List, err error) {
	var l []models.List
	result := repo.DB.Where("owner = ? and active = ?", owner, active).Find(&lists)
	if result.Error != nil {
		return l, nil
	} else {
		return l, result.Error
	}
}

func (repo ListNotesRepositoryImpl) Delete(listId uint) {
	var list models.List
	result := repo.DB.Where("id = ?", listId).Delete(&list)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func (repo ListNotesRepositoryImpl) Disable(listId uint) {
	var list models.List
	result := repo.DB.Model(list).Where("id = ?", listId).Update("active", false)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
