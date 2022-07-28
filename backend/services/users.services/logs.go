package usersservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RetrieveAllLogs() []models.Notifications {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var logs []models.Notifications

	result := db.Find(&logs)

	if result.Error != nil {
		panic(result.Error)
	}

	return logs
}
