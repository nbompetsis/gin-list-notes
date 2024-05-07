package service

type ListNotesService interface {
	Create(serviceList ListData) error
	Update(listData ListData) error
	AddNotesToList(listData ListData) error
}
