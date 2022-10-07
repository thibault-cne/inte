package suggestions

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func All(ctx *gin.Context) {
	suggestions, err := models.RetrieveAllSuggestions()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, suggestions)
}
