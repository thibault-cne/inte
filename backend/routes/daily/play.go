package daily

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Play(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	result, err := models.PlayDailyGame(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "result": result})
}
