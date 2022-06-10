package notificationservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ReadNotification(id int) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	var notification models.Notifications

	db.First(&notification, id)

	notification.Read = true

	db.Save(&notification)

	return nil
}

func RetriveAllUserNotification(user_id int) ([]models.Notifications, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var notifications []models.Notifications

	db.Where("user_id = ? AND read = false", user_id).Find(&notifications)

	return notifications, nil
}
