package dailygameservices

import (
	"backend/models"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
