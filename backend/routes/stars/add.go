package stars

import (
	"backend/models"
	stars_services "backend/services/stars.services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	user := ctx.MustGet("User").(*models.User)

	if user.Current_year == 1 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User must be in year 2 to add stars"})
	}

	targetId := ctx.PostForm("user_id")

	star_type, err := strconv.Atoi(ctx.PostForm("star_type"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid star_type"})
	}

	message := ctx.PostForm("message")

	err = stars_services.AddStars(stars_services.NewStars(user.ID, targetId, star_type, message))

	if err != nil {
		if err.Error() == "message size" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "message size"})
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
