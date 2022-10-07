package stars

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	stars, err := models.GetAllStars()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, models.NewStarsResponse(stars))
}
