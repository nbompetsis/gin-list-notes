package service

type ListData struct {
	ID     uint
	Name   string
	Owner  string
	Active bool
	Notes  []NoteData
}

type NoteData struct {
	ID      uint
	Name    string
	Checked bool
}
