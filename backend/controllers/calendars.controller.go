package controllers

import (
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Return a list of all calendars with the day <= the day parameter
func check_calendars(ctx *gin.Context) {
	day, err := strconv.Atoi(ctx.Param("day"))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Internal server error"})
		return
	}

	calendars, err := services.GetAllCalendars(day)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": calendars})
}

func Register_calendars_routes(rg *gin.RouterGroup) {
	router_group := rg.Group("/calendars")
	router_group.GET("/check/:day", check_calendars)
}
