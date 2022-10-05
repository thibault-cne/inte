package daily

import (
	daily_game_services "backend/services/daily_game.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Play(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	result, err := daily_game_services.PlayDailyGame(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "result": result})
}
