package starsservices

import (
	"backend/models"
	notifications_services "backend/services/notification.services"
	users_services "backend/services/users.services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

		giver, err := users_services.GetUser(star.Giver_id)

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

		notification := notifications_services.NewNotification(star.Receiver_id, "star", "Vous avez reçus une étoile de "+message+" de la part de "+giver.Name)
		notifications_services.AddNewNotification(notification)
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

	db.Where("receiver_id = ? AND type = ? AND moderation_status = 1", user_id, type_).Find(&stars)

	return len(stars), nil
}
