package calendar

import (
	"backend/models"
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

	calendar, err := models.GetCalendarById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Incorrect id"})
	}

	day := ctx.PostForm("day")
	date := ctx.PostForm("date")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")

	models.ModifyCalendar(calendar, date, day, title, content)
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
