package models

import "gorm.io/gorm"

type Challenge struct {
	gorm.Model
	ID          int    `json:"id"`
	Number      int    `json:"number"`
	Title       string `json:"title"`
	Description string `json:"description"`
	WinnerId   int    `json:"winner_id"`
	BeginDate  string `json:"begin_date"`
	EndDate    string `json:"end_date"`
	Status      bool   `json:"status"`
}
