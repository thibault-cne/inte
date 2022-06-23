package planningservices

import (
	"backend/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RetrieveLastPlanning() (*models.Planning, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	today := time.Now().Format("2006-01-02")

	var planning models.Planning

	db.Where("Spawn_time <= ? AND End_time >= ?", today, today).Last(&planning)

	return &planning, nil
}
