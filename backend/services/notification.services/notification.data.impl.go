package notificationservices

import (
	"backend/db"
	"backend/models"
)

func ReadNotification(id int) error {
	var notification *models.Notifications

	db.DB.First(&notification, id)

	notification.Read = true

	db.DB.Save(&notification)

	return nil
}

func RetrieveAllUserNotification(user_id int) ([]*models.Notifications, error) {
	var notifications []*models.Notifications

	db.DB.Where("user_id = ? AND read = false", user_id).Find(&notifications)

	return notifications, nil
}
