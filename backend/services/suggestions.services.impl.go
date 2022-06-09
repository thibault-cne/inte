package services

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSuggestions(title string, description string, user_id int) *models.Suggestion {
	return &models.Suggestion{Title: title, Content: description, Giver_id: user_id}
}

func RegisterSuggestions(suggestion *models.Suggestion) (int, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return 0, err
	}

	db.Create(suggestion)

	return suggestion.ID, nil
}

func RetrieveAllSuggestions() ([]models.Suggestion, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var suggestions []models.Suggestion

	db.Find(&suggestions)

	return suggestions, nil
}
