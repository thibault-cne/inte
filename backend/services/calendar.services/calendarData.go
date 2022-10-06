package calendarservices

import (
	"backend/db"
	"backend/models"
	"strconv"
)

func ModifyCalendar(calendar *models.Calendar, date string, day string, title string, content string) {
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
