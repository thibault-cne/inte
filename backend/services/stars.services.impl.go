package services

import (
	"backend/models"
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewStars(giver_id int, receiver_id int, type_ int, message string) models.Stars {
	return models.Stars{
		Giver_id:                  giver_id,
		Receiver_id:               receiver_id,
		Type:                      type_,
		Message:                   message,
		Moderation_pending_status: 0,
		Moderation_status:         false,
	}
}

func GetAllStars() ([]models.Stars, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var stars []models.Stars

	db.Where("moderation_status = true").Find(&stars)

	return stars, nil
}

func AddStars(star models.Stars) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	if len(star.Message) > 100 || len(star.Message) == 0 {
		err = errors.New("message size")
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

	db.Where("receiver_id = ?, moderation_status = true", user_id).Find(&stars)

	return stars, err
}

// Function to moderate a star
func ModerateStar(id int, user_id int) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	var star models.Stars

	db.First(&star, id)

	if star.Moderation_pending_status != 0 && star.Moderation_pending_status != user_id {
		star.Moderation_status = true

		giver, err := GetUser(star.Giver_id)

		if err != nil {
			return err
		}

		message := ""

		if star.Type == 0 {
			message = "bronze"
		} else if star.Type == 1 {
			message = "silver"
		} else if star.Type == 2 {
			message = "gold"
		}

		notification := NewNotification(star.Receiver_id, "star", "Vous avez reçus une étoile de "+message+" de la part de "+giver.Name)
		RegisterNewNotification(notification)
	} else {
		star.Moderation_pending_status = user_id
	}

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
