package models

import (
	"backend/db"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type DailyGame struct {
	gorm.Model
	ID         int    `json:"id"`
	UserId    int    `json:"user_id"`
	Result     int    `json:"result"`
	CreatedAt string `json:"created_at"`
}

func NewDailyGame(user_id int, result int) *DailyGame {
	return &DailyGame{UserId: user_id, Result: result, CreatedAt: time.Now().Format("2006-01-02")}
}

func CheckDailyGame(user_id int) (bool, error) {
	var daily_game *DailyGame

	// Format date to format 2022-05-19 15:33:00.040402+02:00
	today := time.Now().Format("2006-01-02")
	// Check if their is a row in the daily_game table with the user_id and today's date
	if err := db.DB.Where("user_id = ? AND created_at = ?", user_id, today).First(&daily_game).Error; err != nil {
		return false, nil
	}

	return true, nil
}

func PlayDailyGame(user_id int) (int, error) {
	// Get a random number between -2 and 3
	random_number := rand.Intn(5) - 2

	// Add a row to the daily_game table with the user_id and today's date
	daily_game := NewDailyGame(user_id, random_number)
	db.DB.Create(&daily_game)

	return random_number, nil
}
