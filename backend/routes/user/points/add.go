package points

import (
	"backend/models"
	users_services "backend/services/users.services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	user := ctx.MustGet("User").(*models.User)

	targetID := ctx.PostForm("user_id")

	points, err := strconv.Atoi(ctx.PostForm("points"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid points"})
	}

	err = users_services.AddPoints(user.ID, targetID, points)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
