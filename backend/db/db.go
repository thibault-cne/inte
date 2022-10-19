package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// Init env variables
	// config.InitEnv()

	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"))
	DB.Statement.RaiseErrorOnNotFound = false

	if err != nil {
		panic(err)
	}
}
