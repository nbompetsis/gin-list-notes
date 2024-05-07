package service

import (
	"github.com/nbompetsis/gin-list-notes/app/models"
	"github.com/nbompetsis/gin-list-notes/app/repository"
)

type ListNotesServiceImpl struct {
	repository repository.ListNotesRepository
}

func NewListNotesServiceImpl(repo repository.ListNotesRepository) ListNotesService {
	return ListNotesServiceImpl{
		repository: repo,
	}
}

func (s ListNotesServiceImpl) Create(listData ListData) error {
	list := mapToList(listData)
	err := s.repository.Save(list)
	if err != nil {
		return err
	}
	return nil
}

func (s ListNotesServiceImpl) Update(listData ListData) error {
	list := models.List{
		Name:   listData.Name,
		Owner:  listData.Owner,
		Active: listData.Active,
	}
	err := s.repository.Update(listData.ID, list)
	if err != nil {
		return err
	}
	return nil
}

func (s ListNotesServiceImpl) AddNotesToList(listData ListData) error {
	list := mapToList(listData)
	err := s.repository.AddListNotes(list.ID, list.Notes)
	if err != nil {
		return err
	}
	return nil
}

func mapToList(listData ListData) models.List {
	list := models.List{
		ID:     listData.ID,
		Name:   listData.Name,
		Owner:  listData.Owner,
		Active: listData.Active,
	}
	var notes []models.Note
	for _, n := range listData.Notes {
		notes = append(notes, models.Note{Name: n.Name})
	}
	list.Notes = notes
	return list
}
