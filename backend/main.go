package main

import (
	"backend/db"
	"backend/server"
	"os"
)

func main() {
	db.InitDatabase()
	server.InitServer()

	POPULATE_TEST_DATABASE := os.Getenv("POPULATE_TEST_DATABASE")

	if POPULATE_TEST_DATABASE == "true" {
		PopulateDefault()
	}
}
