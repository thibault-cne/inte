package planningservices

import (
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RetrieveLastPlanning() (*Planning, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	today := time.Now().Format("2006-01-02")

	var planning Planning

	db.Where("Spawn_time <= ? AND End_time >= ?", today, today).Last(&planning)

	return &planning, nil
}

func RetrievePlanningById(planningId string) *Planning {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil
	}

	var planning *Planning

	id, err := strconv.Atoi(planningId)

	if err != nil {
		return nil
	}

	db.First(planning, id)
	return planning
}
