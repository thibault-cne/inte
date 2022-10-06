package dailygameservices

import (
	"backend/db"
	"backend/models"
	"fmt"
	"time"
)

func CheckDailyGame(user_id int) (bool, error) {
	var daily_game models.DailyGame

	// Format date to format 2022-05-19 15:33:00.040402+02:00
	today := time.Now().Format("2006-01-02")
	fmt.Println(today)
	// Check if their is a row in the daily_game table with the user_id and today's date
	if err := db.DB.Where("user_id = ? AND created_at = ?", user_id, today).First(&daily_game).Error; err != nil {
		return false, nil
	}

	return true, nil
}
