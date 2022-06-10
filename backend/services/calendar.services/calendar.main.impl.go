package calendarservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AddCalendar(calendar models.Calendar) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	if err := db.Create(&calendar).Error; err != nil {
		return err
	}

	return nil
}
