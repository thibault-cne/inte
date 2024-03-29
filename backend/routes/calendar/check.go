package calendar

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Return a list of all calendars with the day <= the day parameter
func Check(ctx *gin.Context) {
	day, err := strconv.Atoi(ctx.Param("day"))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect day"})
		return
	}

	calendars, err := models.GetAllCalendars(day)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": calendars})
}
