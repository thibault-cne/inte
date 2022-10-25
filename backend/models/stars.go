package models

import (
	"backend/db"
	"errors"

	"gorm.io/gorm"
)

type Stars struct {
	gorm.Model
	GiverId                  string `json:"giver_id"`
	ReceiverId               string `json:"receiver_id"`
	Type                      int    `json:"type"`
	Message                   string `json:"message"`
	ModerationPendingStatus string `json:"moderation_pending_status"`
	ModerationStatus         bool   `json:"moderation_status"`
}

func NewStars(giver_id string, receiver_id string, type_ int, message string) *Stars {
	return &Stars{
		GiverId:          giver_id,
		ReceiverId:       receiver_id,
		Type:              type_,
		Message:           message,
		ModerationStatus: false,
	}
}

func AddStars(star *Stars) error {
	if len(star.Message) > 100 || len(star.Message) == 0 {
		return errors.New("message size")
	}

	db.DB.Create(&star)

	return nil
}

// Function to moderate a star
func ModerateStar(id int, user_id string) error {
	var star Stars

	db.DB.First(&star, id)

	if star.ModerationStatus {
		return nil
	}

	if star.ModerationPendingStatus != "" && star.ModerationPendingStatus != user_id {
		star.ModerationStatus = true

		giver, err := GetUser(star.GiverId)

		if err != nil {
			return err
		}

		message := ""
		points := 0

		if star.Type == 0 {
			message = "bronze"
			points = 3
		} else if star.Type == 1 {
			message = "silver"
			points = 5
		} else if star.Type == 2 {
			message = "gold"
			points = 7
		}

		notification := NewNotification(star.ReceiverId, "star", "Vous avez reçus une étoile de "+message+" de la part de "+giver.Name)
		AddNewNotification(notification)
		AddPoints(star.GiverId, star.ReceiverId, points)
	} else {
		star.ModerationPendingStatus = user_id
	}

	db.DB.Save(&star)

	return nil
}

func CountStarsType(user_id string, type_ int) (int, error) {
	var stars []*Stars

	db.DB.Where("receiver_id = ? AND type = ? AND moderation_status = 1", user_id, type_).Find(&stars)

	return len(stars), nil
}

func GetAllStars() ([]*Stars, error) {
	var stars []*Stars

	db.DB.Where("moderation_status = true").Find(&stars)

	return stars, nil
}

func GetStarsPage(offset int, amount int) ([]*Stars, error) {
	var stars []*Stars

	e := db.DB.Where("moderation_status = true").Offset(offset).Limit(amount).Find(&stars)

	return stars, e.Error
}

// Get all stars of a user
func GetStars(user_id string) ([]*Stars, error) {
	var stars []*Stars

	e := db.DB.Where("receiver_id = ?, moderation_status = true", user_id).Find(&stars)

	return stars, e.Error
}
