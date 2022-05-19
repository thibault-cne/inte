package models

import "gorm.io/gorm"

type Calendar struct {
	gorm.Model
	ID      int    `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
