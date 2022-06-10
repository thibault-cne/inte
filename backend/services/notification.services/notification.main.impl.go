package notificationservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewNotification(user_id int, type_ string, message string) models.Notifications {
	return models.Notifications{
		User_id: user_id,
		Type:    type_,
		Message: message,
		Read:    false,
	}
}

func AddNewNotification(notification models.Notifications) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Create(&notification)

	return nil
}
