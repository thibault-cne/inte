package models

import "gorm.io/gorm"

type Suggestion struct {
	gorm.Model
	ID       int    `json:"id"`
	Giver_id int    `json:"giver_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
