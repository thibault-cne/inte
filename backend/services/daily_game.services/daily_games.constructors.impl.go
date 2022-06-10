package dailygameservices

import (
	"backend/models"
	"time"
)

func NewDailyGame(user_id int, result int) models.DailyGame {
	return models.DailyGame{User_id: user_id, Result: result, Created_at: time.Now().Format("2006-01-02")}
}
