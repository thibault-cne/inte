package starsservices

import (
	"backend/models"
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewStars(giver_id string, receiver_id string, type_ int, message string) models.Stars {
	return models.Stars{
		Giver_id:          giver_id,
		Receiver_id:       receiver_id,
		Type:              type_,
		Message:           message,
		Moderation_status: false,
	}
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
