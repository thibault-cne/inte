package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	// init environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
