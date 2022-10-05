package usersservices

import (
	"backend/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewUser(email string, name string) *models.User {
	t := time.Now()

	year := 0

	if t.Month() < 9 {
		year = t.Year() - 1 + 3
	} else {
		year = t.Year() + 3
	}

	return &models.User{Name: name, Email: email, Current_year: 1, Promotion_year: year, Points: 0, User_type: "user"}
}

func AddUser(user *models.User) (string, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return "", err
	}

	db.Create(user)

	return user.ID, nil
}
