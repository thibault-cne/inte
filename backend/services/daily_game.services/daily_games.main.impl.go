package dailygameservices

import (
	"backend/db"
	"math/rand"
)

func PlayDailyGame(user_id int) (int, error) {
	// Get a random number between -2 and 3
	random_number := rand.Intn(5) - 2

	// Add a row to the daily_game table with the user_id and today's date
	daily_game := NewDailyGame(user_id, random_number)
	db.DB.Create(&daily_game)

	return random_number, nil
}
