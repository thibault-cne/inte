package models

import "gorm.io/gorm"

type Challenge struct {
	gorm.Model
	ID          int    `json:"id"`
	Number      int    `json:"number"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Winner_id   int    `json:"winner_id"`
	Begin_date  string `json:"begin_date"`
	End_date    string `json:"end_date"`
	Status      bool   `json:"status"`
}
