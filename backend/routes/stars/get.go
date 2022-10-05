package stars

import (
	api_services "backend/services/api_response.services"
	stars_services "backend/services/stars.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	stars, err := stars_services.GetAllStars()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, api_services.NewStarsResponse(stars))
}
