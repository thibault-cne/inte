package newsinteservices

import (
	"backend/db"
	"backend/models"
)

func NewNewsInte(content string) *models.NewsInte {
	return &models.NewsInte{Content: content}
}

func AddNewsInte(news *models.NewsInte) {
	db.DB.Create(news)
}

func EditNews(id int, newContent string) {
	var news *models.NewsInte

	result := db.DB.Where("id = ?", id).Find(&news)

	if result.RowsAffected == 0 {
		return
	}

	news.Content = newContent

	db.DB.Save(news)
}

func DeleteNews(id int) {
	db.DB.Delete(&models.NewsInte{}, id)
}
