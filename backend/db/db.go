package db

import (
	"backend/config"
	"backend/models"
	planningservices "backend/services/planning.services"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() {
	// Init env variables
	config.InitEnv()

	POPULATE_TEST_DATABASE := os.Getenv("POPULATE_TEST_DATABASE")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
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
	db.AutoMigrate(&planningservices.Planning{})

	if POPULATE_TEST_DATABASE == "true" {
		populateTestDatabase()
	}
}
