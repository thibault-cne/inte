package planningservices

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewPlaning(picture string, spawn_time string, end_time string) Planning {
	return Planning{
		Picture:    picture,
		Spawn_time: spawn_time,
		End_time:   end_time,
	}
}

func AddPlaning(planing Planning) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Create(&planing)

	return nil
}
