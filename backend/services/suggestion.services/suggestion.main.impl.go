package suggestionservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSuggestions(title string, description string, user_id int) *models.Suggestion {
	return &models.Suggestion{Title: title, Content: description, Giver_id: user_id}
}

func AddSuggestions(suggestion *models.Suggestion) (int, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return 0, err
	}

	db.Create(suggestion)

	return suggestion.ID, nil
}
