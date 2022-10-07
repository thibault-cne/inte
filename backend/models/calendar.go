package models

import (
	"backend/db"
	"strconv"

	"gorm.io/gorm"
)

type Calendar struct {
	gorm.Model
	ID      int    `json:"id"`
	Date    string `json:"date"`
	Day     int    `json:"day"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func AddCalendar(calendar *Calendar) error {
	if err := db.DB.Create(&calendar).Error; err != nil {
		return err
	}

	return nil
}

func GetAllCalendars(day int) ([]*Calendar, error) {
	var calendars []*Calendar

	if err := db.DB.Where("day <= ?", day).Find(&calendars).Error; err != nil {
		return nil, err
	}

	return calendars, nil
}

func GetCalendarById(id int) (*Calendar, error) {
	var calendars *Calendar

	if err := db.DB.Where("id = ?", id).Find(&calendars).Error; err != nil {
		return nil, err
	}

	return calendars, nil
}

func ModifyCalendar(calendar *Calendar, date string, day string, title string, content string) {
	if date != "" {
		calendar.Date = date
	}

	if day != "" {
		dayInteger, err := strconv.Atoi(day)

		if err != nil {
			panic(err)
		}

		calendar.Day = dayInteger
	}

	if title != "" {
		calendar.Title = title
	}

	if content != "" {
		calendar.Content = content
	}

	db.DB.Save(calendar)
}

func (c *Calendar) Save() {
	db.DB.Save(c)
}
