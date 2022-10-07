package get

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Notifs(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	users, err := models.RetrieveAllUserNotification(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
