package controllers

import (
	"backend/models"
	calendar_services "backend/services/calendar.services"
	claims_services "backend/services/claims.services"
	users_services "backend/services/users.services"
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
	reqToken := ctx.Request.Header.Get("Authorization")
	claims, err := claims_services.RetrieveUserClaims(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	admin_status, err := users_services.CheckAdmin(claims.User_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if !admin_status {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not an admin"})
		return
	}

	var calendar models.Calendar

	err = ctx.BindJSON(&calendar)

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

func Register_calendars_routes(rg *gin.RouterGroup) {
	router_group := rg.Group("/calendars")
	router_group.POST("/add/day", add_calendar)
	router_group.GET("/check/:day", check_calendars)
}
