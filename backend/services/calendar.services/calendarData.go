package calendarservices

import (
	"backend/models"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ModifyCalendar(calendar *models.Calendar, date string, day string, title string, content string) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

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

	db.Save(calendar)
}
