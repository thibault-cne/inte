package get

import (
	"backend/models"
	stars_services "backend/services/stars.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Stats(ctx *gin.Context) {
	user := ctx.MustGet("User").(*models.User)

	gold, err := stars_services.CountStarsType(user.ID, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	silver, err := stars_services.CountStarsType(user.ID, 1)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	bronze, err := stars_services.CountStarsType(user.ID, 2)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"points": user.Points, "gold_stars": gold, "silver_stars": silver, "bronze_stars": bronze})
}
