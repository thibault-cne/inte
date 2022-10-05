package stars

import (
	"net/http"
	"strconv"

	stars_services "backend/services/stars.services"
	users_services "backend/services/users.services"

	"github.com/gin-gonic/gin"
)

func Moderate(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	// Verify if user is admin
	user, err := users_services.GetUser(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	if user.User_type != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User is not admin"})
	}

	star_id, err := strconv.Atoi(ctx.PostForm("star_id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid star_id"})
	}

	err = stars_services.ModerateStar(star_id, userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
