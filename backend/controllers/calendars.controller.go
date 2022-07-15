package controllers

import (
	"backend/models"
	calendar_services "backend/services/calendar.services"
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

	calendars, err := calendar_services.GetAllCalendars(day)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": calendars})
}

func add_calendar(ctx *gin.Context) {
	var calendar models.Calendar

	err := ctx.BindJSON(&calendar)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Internal server error"})
		return
	}

	err = calendar_services.AddCalendar(calendar)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func registerCalendarsRoutes(rg *gin.RouterGroup) {
	router_group := rg.Group("/calendars")
	router_group.GET("/check/:day", check_calendars)
}

func registerAdminCalendarRoutes(rg *gin.RouterGroup) {
	router_group := rg.Group("/calendars")
	router_group.POST("/add/day", add_calendar)
}
