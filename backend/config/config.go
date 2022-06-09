package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func InitEnv() {
	// init environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
