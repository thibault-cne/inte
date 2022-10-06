package calendarservices

import (
	"backend/db"
	"backend/models"
)

func GetAllCalendars(day int) ([]models.Calendar, error) {
	var calendars []models.Calendar

	if err := db.DB.Where("day <= ?", day).Find(&calendars).Error; err != nil {
		return nil, err
	}

	return calendars, nil
}

func GetCalendarById(id int) (*models.Calendar, error) {
	var calendars *models.Calendar

	if err := db.DB.Where("id = ?", id).Find(&calendars).Error; err != nil {
		return nil, err
	}

	return calendars, nil
}
