package services

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

func RegisterNewNotification(notification models.Notifications) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Create(&notification)

	return nil
}

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
