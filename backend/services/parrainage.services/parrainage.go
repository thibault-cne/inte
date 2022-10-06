package parrainageservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RetrieveCurrentParrainage(userId string) *models.Parrainage {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var parrainage models.Parrainage

	result := db.Where("god_father_id = ?", userId).Find(&parrainage)

	if result.RowsAffected == 0 {
		newParrainage := createParrainage(userId)
		newParrainage.AddParrainage()

		return newParrainage
	}

	db.Where("godFatherId = ?", userId).First(parrainage)

	return &parrainage
}
