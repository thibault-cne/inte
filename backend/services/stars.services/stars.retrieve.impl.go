package starsservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetAllStars() ([]models.Stars, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var stars []models.Stars

	db.Where("moderation_status = true").Find(&stars)

	return stars, nil
}

// Get all stars of a user
func GetStars(user_id int) ([]models.Stars, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var stars []models.Stars

	db.Where("receiver_id = ?, moderation_status = true", user_id).Find(&stars)

	return stars, err
}
