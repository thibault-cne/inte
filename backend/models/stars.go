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

func (s *Stars) Create() error {
	if len(s.Message) > 100 || len(s.Message) == 0 {
		return errors.New("message size")
	}

	if err := db.DB.Create(s).Error; err != nil {
		return err
	}

	return nil
}

// Function to moderate a star
func (u *User) ModerateStar(star *Stars) error {
	if star.ModerationStatus {
		return nil
	}

	if star.ModerationPendingStatus != "" && star.ModerationPendingStatus != u.ID {
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
		notification.Create()
		AddPoints(star.GiverId, star.ReceiverId, points)
	} else {
		star.ModerationPendingStatus = u.ID
	}

	db.DB.Save(&star)

	return nil
}

func GetStar(star_id uint) *Stars {
	var star *Stars

	db.DB.Where("ID = ?", star_id).Find(star)

	return star
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
func (u *User) GetStars() {
	e := db.DB.Where("receiver_id = ? AND moderation_status = true", u.ID).Find(&u.Stars)

	if e.Error != nil {
		panic(e.Error)
	}
}
