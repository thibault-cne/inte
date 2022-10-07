package planning

import (
	"backend/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Current(ctx *gin.Context) {
	currentPlanning, err := models.RetrieveLastPlanning()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	// planningResponse, err := api_response_services.NewPlanningResponse(currentPlanning)

	fileContent, err := os.ReadFile("static/images/planning_pictures/" + currentPlanning.Picture)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.Data(http.StatusOK, "image/jpeg", fileContent)
}
