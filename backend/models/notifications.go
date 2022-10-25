package models

import (
	"backend/db"

	"gorm.io/gorm"
)

type Notifications struct {
	gorm.Model
	ID      int    `json:"id"`
	UserId string `json:"user_id"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Read    bool   `json:"read"`
}

func ReadNotification(id int) error {
	var notification *Notifications

	db.DB.First(&notification, id)

	notification.Read = true

	db.DB.Save(&notification)

	return nil
}

func RetrieveAllUserNotification(user_id int) ([]*Notifications, error) {
	var notifications []*Notifications

	db.DB.Where("user_id = ? AND read = false", user_id).Find(&notifications)

	return notifications, nil
}

func NewNotification(user_id string, type_ string, message string) *Notifications {
	return &Notifications{
		UserId: user_id,
		Type:    type_,
		Message: message,
		Read:    false,
	}
}

func AddNewNotification(notification *Notifications) error {
	db.DB.Create(&notification)
	return nil
}
