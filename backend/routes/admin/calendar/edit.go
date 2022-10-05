package calendar

import (
	calendar_services "backend/services/calendar.services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Edit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect id"})
		return
	}

	calendar, err := calendar_services.GetCalendarById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Incorrect id"})
	}

	day := ctx.PostForm("day")
	date := ctx.PostForm("date")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")

	calendar_services.ModifyCalendar(calendar, date, day, title, content)
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
