package calendar

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddDay(ctx *gin.Context) {
	var calendar *models.Calendar

	err := ctx.BindJSON(&calendar)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Internal server error"})
	}

	err = calendar.Create()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
