package main

import (
	"backend/db"
	"backend/models"
	"backend/server"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db.InitDatabase()
	models.Migrate()

	POPULATE_TEST_DATABASE := os.Getenv("POPULATE_TEST_DATABASE")
	if POPULATE_TEST_DATABASE == "true" {
		PopulateDefault()
	}

	server.InitServer()
}
