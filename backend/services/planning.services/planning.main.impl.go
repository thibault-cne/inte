package planningservices

import (
	"backend/db"
	"backend/models"
)

func NewPlaning(picture string, spawn_time string, end_time string) *models.Planning {
	return &models.Planning{
		Picture:    picture,
		Spawn_time: spawn_time,
		End_time:   end_time,
	}
}

func AddPlaning(planing *models.Planning) error {
	db.DB.Create(&planing)
	return nil
}
