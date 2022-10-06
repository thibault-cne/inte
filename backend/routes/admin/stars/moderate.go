package stars

import (
	"net/http"
	"strconv"

	"backend/models"
	stars_services "backend/services/stars.services"

	"github.com/gin-gonic/gin"
)

func Moderate(ctx *gin.Context) {
	user := ctx.MustGet("User").(*models.User)

	star_id, err := strconv.Atoi(ctx.PostForm("star_id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid star_id"})
	}

	err = stars_services.ModerateStar(star_id, user.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
