package stars

import (
	stars_services "backend/services/stars.services"
	users_services "backend/services/users.services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	// Verify if user is in year > 2
	user, err := users_services.GetUser(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	if user.Current_year == 1 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User must be in year 2 to add stars"})
	}

	user_id, err := strconv.Atoi(ctx.PostForm("user_id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
	}

	star_type, err := strconv.Atoi(ctx.PostForm("star_type"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid star_type"})
	}

	message := ctx.PostForm("message")

	err = stars_services.AddStars(stars_services.NewStars(userId, user_id, star_type, message))

	if err != nil {
		if err.Error() == "message size" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "message size"})
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
