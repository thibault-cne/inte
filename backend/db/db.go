package db

import (
	"backend/config"
	"backend/models"
	parrainageservices "backend/services/parrainage.services"
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

	if POPULATE_TEST_DATABASE == "true" {
		db.Migrator().DropTable("users")
		db.Migrator().DropTable("stars")
		db.Migrator().DropTable("plannings")
		db.Migrator().DropTable("suggestions")

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
		db.AutoMigrate(&parrainageservices.Parrainage{})
		db.AutoMigrate(&parrainageservices.Adoption{})

		populateTestDatabase()

	} else {
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
		db.AutoMigrate(&parrainageservices.Parrainage{})
		db.AutoMigrate(&parrainageservices.Adoption{})
	}
}
