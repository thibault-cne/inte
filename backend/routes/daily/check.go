package daily

import (
	daily_game_services "backend/services/daily_game.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Check(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	status, err := daily_game_services.CheckDailyGame(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": status})
}
