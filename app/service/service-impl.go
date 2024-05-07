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

func (s ListNotesServiceImpl) FindListNotesByListID(listID uint) (ListData, error) {
	listNotesInfo, err := s.repository.FindListNotesByListID(listID)
	if err != nil {
		return ListData{}, err
	}
	return mapToListData([]models.ListNotesInfo{listNotesInfo}), nil
}

func (s ListNotesServiceImpl) FindListNotesByOwner(owner string) (ListData, error) {
	listNotesInfo, err := s.repository.FindListNotesByOwner(owner)
	if err != nil {
		return ListData{}, err
	}
	return mapToListData(listNotesInfo), nil
}

func (s ListNotesServiceImpl) DeleteList(listID uint) error {
	err := s.repository.DeleteList(listID)
	if err != nil {
		return err
	}
	return nil
}

func (s ListNotesServiceImpl) CheckListNote(listID uint, noteID uint) error {
	err := s.repository.CheckListNote(listID, noteID)
	if err != nil {
		return err
	}
	return nil
}

func (s ListNotesServiceImpl) CheckListAllNotes(listID uint) error {
	err := s.repository.CheckListAllNotes(listID)
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

func mapToListData(listNotesInfo []models.ListNotesInfo) ListData {
	if len(listNotesInfo) == 0 {
		return ListData{}
	}

	var notes []NoteData
	for _, n := range listNotesInfo {
		notes = append(notes, NoteData{Name: n.NoteName, Checked: n.NoteChecked})
	}
	l := ListData{
		ID:     listNotesInfo[0].ListID,
		Name:   listNotesInfo[0].ListName,
		Active: listNotesInfo[0].ListActive,
		Notes:  notes,
	}
	return l
}
