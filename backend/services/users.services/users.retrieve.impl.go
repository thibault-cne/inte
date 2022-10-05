package usersservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetUser(id string) (*models.User, error) {
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

func GetUserByName(name string) *models.User {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var user models.User

	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		panic(err)
	}

	return &user
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
