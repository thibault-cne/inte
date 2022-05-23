package main

import (
	"backend/models"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestDatabaseInit(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test_database.db"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Stars{})
	db.AutoMigrate(&models.DailyGame{})
	db.AutoMigrate(&models.Notifications{})
	db.AutoMigrate(&models.Calendar{})
	db.AutoMigrate(&models.Challenge{})
	db.AutoMigrate(&models.Suggestion{})
	db.AutoMigrate(&models.Tnder{})

	status := db.Migrator().HasTable(&models.User{})

	if !status {
		t.Error("Table User not found")
	}

	status = db.Migrator().HasTable(&models.Stars{})

	if !status {
		t.Error("Table Stars not found")
	}

	status = db.Migrator().HasTable(&models.DailyGame{})

	if !status {
		t.Error("Table DailyGame not found")
	}

	status = db.Migrator().HasTable(&models.Notifications{})

	if !status {
		t.Error("Table Notifications not found")
	}

	status = db.Migrator().HasTable(&models.Calendar{})

	if !status {
		t.Error("Table Calendar not found")
	}

	status = db.Migrator().HasTable(&models.Challenge{})

	if !status {
		t.Error("Table Challenge not found")
	}

	status = db.Migrator().HasTable(&models.Suggestion{})

	if !status {
		t.Error("Table Suggestion not found")
	}

	status = db.Migrator().HasTable(&models.Tnder{})

	if !status {
		t.Error("Table Tnder not found")
	}
}
