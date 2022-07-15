package userscontrollers

import (
	users_services "backend/services/users.services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func add_points(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	user_id, err := strconv.Atoi(ctx.PostForm("user_id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	points, err := strconv.Atoi(ctx.PostForm("points"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid points"})
		return
	}

	err = users_services.AddPoints(userId, user_id, points)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func Register_user_controllers_points(rg *gin.RouterGroup) {
	router_group := rg.Group("/points")
	router_group.POST("/add", add_points)
}
