package suggestionservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RetrieveAllSuggestions() ([]models.Suggestion, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var suggestions []models.Suggestion

	db.Find(&suggestions)

	return suggestions, nil
}
