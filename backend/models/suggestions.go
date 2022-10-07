package models

import (
	"backend/db"

	"gorm.io/gorm"
)

type Suggestion struct {
	gorm.Model
	ID       int    `json:"id"`
	Giver_id int    `json:"giver_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func NewSuggestions(title string, description string, user_id int) *Suggestion {
	return &Suggestion{Title: title, Content: description, Giver_id: user_id}
}

func (s *Suggestion) Create() (int, error) {
	db.DB.Create(s)

	return s.ID, nil
}

func RetrieveAllSuggestions() ([]*Suggestion, error) {
	var suggestions []*Suggestion

	db.DB.Find(&suggestions)

	return suggestions, nil
}
