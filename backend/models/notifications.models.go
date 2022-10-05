package models

import "gorm.io/gorm"

type Notifications struct {
	gorm.Model
	ID      int    `json:"id"`
	User_id string `json:"user_id"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Read    bool   `json:"read"`
}
