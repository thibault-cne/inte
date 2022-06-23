package models

import "gorm.io/gorm"

type Planning struct {
	gorm.Model
	ID         int    `json:"id"`
	Picture    string `json:"picture"`
	Spawn_time string `json:"spawn_time"`
	End_time   string `json:"end_time"`
}
