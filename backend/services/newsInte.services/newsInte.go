package newsinteservices

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewNewsInte(content string) *NewsInte {
	return &NewsInte{Content: content}
}

func AddNewsInte(news *NewsInte) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Create(news)
}

func EditNews(id int, newContent string) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var news NewsInte

	result := db.Where("id = ?", id).Find(&news)

	if result.RowsAffected == 0 {
		return
	}

	news.Content = newContent

	db.Save(news)
}

func DeleteNews(id int) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Delete(&NewsInte{}, id)
}
