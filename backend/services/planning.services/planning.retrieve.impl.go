package planningservices

import (
	"backend/db"
	"backend/models"
	"strconv"
	"time"
)

func RetrieveLastPlanning() (*models.Planning, error) {
	today := time.Now().Format("2006-01-02")

	var planning *models.Planning

	db.DB.Where("Spawn_time <= ? AND End_time >= ?", today, today).Last(&planning)

	return planning, nil
}

func RetrievePlanningById(planningId string) *models.Planning {
	var planning *models.Planning

	id, err := strconv.Atoi(planningId)

	if err != nil {
		return nil
	}

	db.DB.First(planning, id)
	return planning
}
