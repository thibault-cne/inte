package usersservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

// Get all users from the database
func GetAllUsers() ([]models.User, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
