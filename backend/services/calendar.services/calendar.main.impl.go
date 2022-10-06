package calendarservices

import (
	"backend/db"
	"backend/models"
)

func AddCalendar(calendar models.Calendar) error {
	if err := db.DB.Create(&calendar).Error; err != nil {
		return err
	}

	return nil
}
