package services

import (
	"backend/models"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDailyGame(user_id int, result int) models.DailyGame {
	return models.DailyGame{User_id: user_id, Result: result, Created_at: time.Now().Format("2006-01-02")}
}

func CheckDailyGame(user_id int) (bool, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return false, err
	}

	var daily_game models.DailyGame

	// Format date to format 2022-05-19 15:33:00.040402+02:00
	today := time.Now().Format("2006-01-02")
	fmt.Println(today)
	// Check if their is a row in the daily_game table with the user_id and today's date
	if err := db.Where("user_id = ? AND created_at = ?", user_id, today).First(&daily_game).Error; err != nil {
		return false, nil
	}

	return true, nil
}

func PlayDailyGame(user_id int) (int, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return 0, err
	}

	// Get a random number between -2 and 3
	random_number := rand.Intn(5) - 2

	// Add a row to the daily_game table with the user_id and today's date
	daily_game := NewDailyGame(user_id, random_number)
	db.Create(&daily_game)

	return random_number, nil
}
