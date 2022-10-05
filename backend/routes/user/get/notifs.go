package get

import (
	notifications_services "backend/services/notification.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Notifs(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	users, err := notifications_services.RetriveAllUserNotification(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
