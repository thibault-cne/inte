package newsinteservices

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RetrieveAllNews() []NewsInte {
	var news []NewsInte

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	result := db.Find(&news)

	if result.Error != nil {
		panic(err)
	}

	return news
}
