package newsinteservices

import "gorm.io/gorm"

type NewsInte struct {
	gorm.Model
	Content string `json:"content"`
}
