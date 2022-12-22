package models

import (
	"backend/db"

	"gorm.io/gorm"
)

type Suggestion struct {
	gorm.Model
	GiverId int    `json:"giver_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func NewSuggestions(title string, description string, user_id int) *Suggestion {
	return &Suggestion{Title: title, Content: description, GiverId: user_id}
}

func (s *Suggestion) Create() (uint, error) {
	if err := db.DB.Create(s).Error; err != nil {
		return 0, err
	}

	return s.ID, nil
}

func RetrieveAllSuggestions() ([]*Suggestion, error) {
	var suggestions []*Suggestion

	db.DB.Find(&suggestions)

	return suggestions, nil
}
