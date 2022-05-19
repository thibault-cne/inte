package models

import (
	"gorm.io/gorm"
)

type DailyGame struct {
	gorm.Model
	ID         int    `json:"id"`
	User_id    int    `json:"user_id"`
	Result     int    `json:"result"`
	Created_at string `json:"created_at"`
}
