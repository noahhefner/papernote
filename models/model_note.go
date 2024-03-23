package models

type Note struct {
	FileName string `json:"filename"`
	Content  string `json:"content"`
}
