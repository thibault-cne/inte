package models

import (
	"backend/db"

	"gorm.io/gorm"
)

type NewsInte struct {
	gorm.Model
	Content string `json:"content"`
}

func (n *NewsInte) Create() {
	db.DB.Create(n)
}

func NewNewsInte(content string) *NewsInte {
	return &NewsInte{Content: content}
}

func EditNews(id int, newContent string) {
	var news *NewsInte

	result := db.DB.Where("id = ?", id).Find(&news)

	if result.RowsAffected == 0 {
		return
	}

	news.Content = newContent

	db.DB.Save(news)
}

func DeleteNews(id int) {
	db.DB.Delete(&NewsInte{}, id)
}

func RetrieveAllNews() []*NewsInte {
	var news []*NewsInte

	result := db.DB.Find(&news)

	if result.Error != nil {
		panic(result.Error)
	}

	return news
}
