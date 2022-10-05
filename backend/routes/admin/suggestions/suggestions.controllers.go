package suggestions

import (
	suggestion_services "backend/services/suggestion.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func All(ctx *gin.Context) {
	suggestions, err := suggestion_services.RetrieveAllSuggestions()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, suggestions)
}
