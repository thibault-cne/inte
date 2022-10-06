package newsinteservices

import (
	"backend/db"
	"backend/models"
)

func RetrieveAllNews() []*models.NewsInte {
	var news []*models.NewsInte

	result := db.DB.Find(&news)

	if result.Error != nil {
		panic(result.Error)
	}

	return news
}
