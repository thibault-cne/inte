package get

import (
	stars_services "backend/services/stars.services"
	users_services "backend/services/users.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Stats(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	users, err := users_services.GetUser(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	gold, err := stars_services.CountStarsType(userId, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	silver, err := stars_services.CountStarsType(userId, 1)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	bronze, err := stars_services.CountStarsType(userId, 2)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"points": users.Points, "gold_stars": gold, "silver_stars": silver, "bronze_stars": bronze})
}
