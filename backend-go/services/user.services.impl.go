package services

import (
	"backend-go/models"
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

func AddUser(user *models.User) (int, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return 0, err
	}
	db.Create(user)

	return user.ID, nil
}

func GetUser(id int) (*models.User, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var user models.User

	db.First(&user, id)

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var user models.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
