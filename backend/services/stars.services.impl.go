package services

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewStars(giver_id int, receiver_id int, type_ int, message string) models.Stars {
	return models.Stars{
		Giver_id:          giver_id,
		Receiver_id:       receiver_id,
		Type:              type_,
		Message:           message,
		Moderation_status: false,
	}
}

func GetAllStars() ([]models.Stars, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var stars []models.Stars

	db.Find(&stars)

	return stars, nil
}

func AddStars(star models.Stars) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Create(&star)

	return nil
}

// Get all stars of a user
func GetStars(user_id int) ([]models.Stars, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var stars []models.Stars

	db.Where("receiver_id = ?", user_id).Find(&stars)

	return stars, err
}

// Function to moderate a star
func ModerateStar(id int) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	var star models.Stars

	db.First(&star, id)

	star.Moderation_status = true

	db.Save(&star)

	return nil
}

func CountStarsType(user_id int, type_ int) (int, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return 0, err
	}

	var stars []models.Stars

	db.Where("receiver_id = ? AND type = ?", user_id, type_).Find(&stars)

	return len(stars), nil
}
