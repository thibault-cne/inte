package calendarservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetAllCalendars(day int) ([]models.Calendar, error) {
	// Get all calendars with the day <= the day parameter
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var calendars []models.Calendar

	if err := db.Where("day <= ?", day).Find(&calendars).Error; err != nil {
		return nil, err
	}

	return calendars, nil
}
