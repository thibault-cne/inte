package models

import "gorm.io/gorm"

type Notifications struct {
	gorm.Model
	ID      int    `json:"id"`
	User_id int    `json:"user_id"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Read    bool   `json:"read"`
}
