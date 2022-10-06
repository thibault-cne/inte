package models

import (
	"backend/db"
	"os"
)

func init() {
	POPULATE_TEST_DATABASE := os.Getenv("POPULATE_TEST_DATABASE")

	if POPULATE_TEST_DATABASE == "true" {
		db.DB.Migrator().DropTable("users")
		db.DB.Migrator().DropTable("stars")
		db.DB.Migrator().DropTable("plannings")
		db.DB.Migrator().DropTable("suggestions")
		db.DB.Migrator().DropTable("news_inte")

		// Migrate the schema
		db.DB.AutoMigrate(&User{})
		db.DB.AutoMigrate(&Stars{})
		db.DB.AutoMigrate(&DailyGame{})
		db.DB.AutoMigrate(&Notifications{})
		db.DB.AutoMigrate(&Calendar{})
		db.DB.AutoMigrate(&Challenge{})
		db.DB.AutoMigrate(&Suggestion{})
		db.DB.AutoMigrate(&Tnder{})
		db.DB.AutoMigrate(&Planning{})
		db.DB.AutoMigrate(&Parrainage{})
		db.DB.AutoMigrate(&Adoption{})
		db.DB.AutoMigrate(&NewsInte{})
	} else {
		// Migrate the schema
		db.DB.AutoMigrate(&User{})
		db.DB.AutoMigrate(&Stars{})
		db.DB.AutoMigrate(&DailyGame{})
		db.DB.AutoMigrate(&Notifications{})
		db.DB.AutoMigrate(&Calendar{})
		db.DB.AutoMigrate(&Challenge{})
		db.DB.AutoMigrate(&Suggestion{})
		db.DB.AutoMigrate(&Tnder{})
		db.DB.AutoMigrate(&Planning{})
		db.DB.AutoMigrate(&Parrainage{})
		db.DB.AutoMigrate(&Adoption{})
		db.DB.AutoMigrate(&NewsInte{})
	}
}
