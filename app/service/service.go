package service

type ListNotesService interface {
	Create(serviceList ListData) error
	Update(listData ListData) error
	AddNotesToList(listData ListData) error
	FindListNotesByListID(listID uint) (ListData, error)
	FindListNotesByOwner(owner string) (ListData, error)
	DeleteList(listID uint) error
	CheckListNote(listID uint, noteID uint) error
	CheckListAllNotes(listID uint) error
}
