package notificationservices

import (
	"backend/db"
	"backend/models"
)

func NewNotification(user_id string, type_ string, message string) *models.Notifications {
	return &models.Notifications{
		User_id: user_id,
		Type:    type_,
		Message: message,
		Read:    false,
	}
}

func AddNewNotification(notification *models.Notifications) error {
	db.DB.Create(&notification)
	return nil
}
