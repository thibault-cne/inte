package models

import "gorm.io/gorm"

type Calendar struct {
	gorm.Model
	ID      int    `json:"id"`
	Date    string `json:"date"`
	Day     int    `json:"day"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
