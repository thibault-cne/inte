package db

import (
	"backend/config"
	"backend/models"
	newsinteservices "backend/services/newsInte.services"
	parrainageservices "backend/services/parrainage.services"
	planningservices "backend/services/planning.services"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// Init env variables
	config.InitEnv()

	POPULATE_TEST_DATABASE := os.Getenv("POPULATE_TEST_DATABASE")

	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if POPULATE_TEST_DATABASE == "true" {
		DB.Migrator().DropTable("users")
		DB.Migrator().DropTable("stars")
		DB.Migrator().DropTable("plannings")
		DB.Migrator().DropTable("suggestions")
		DB.Migrator().DropTable("news_inte")

		// Migrate the schema
		DB.AutoMigrate(&models.User{})
		DB.AutoMigrate(&models.Stars{})
		DB.AutoMigrate(&models.DailyGame{})
		DB.AutoMigrate(&models.Notifications{})
		DB.AutoMigrate(&models.Calendar{})
		DB.AutoMigrate(&models.Challenge{})
		DB.AutoMigrate(&models.Suggestion{})
		DB.AutoMigrate(&models.Tnder{})
		DB.AutoMigrate(&planningservices.Planning{})
		DB.AutoMigrate(&parrainageservices.Parrainage{})
		DB.AutoMigrate(&parrainageservices.Adoption{})
		DB.AutoMigrate(&newsinteservices.NewsInte{})

		populateTestDatabase()

	} else {
		// Migrate the schema
		DB.AutoMigrate(&models.User{})
		DB.AutoMigrate(&models.Stars{})
		DB.AutoMigrate(&models.DailyGame{})
		DB.AutoMigrate(&models.Notifications{})
		DB.AutoMigrate(&models.Calendar{})
		DB.AutoMigrate(&models.Challenge{})
		DB.AutoMigrate(&models.Suggestion{})
		DB.AutoMigrate(&models.Tnder{})
		DB.AutoMigrate(&planningservices.Planning{})
		DB.AutoMigrate(&parrainageservices.Parrainage{})
		DB.AutoMigrate(&parrainageservices.Adoption{})
		DB.AutoMigrate(&newsinteservices.NewsInte{})
	}
}
