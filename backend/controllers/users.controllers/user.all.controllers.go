package userscontrollers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func get_all_users(ctx *gin.Context) {
	users, err := services.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, services.NewAllUsersResponse(users))
}

func get_all_users_with_points(ctx *gin.Context) {
	users, err := services.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, services.NewAllUsersWithPointsResponse(users))
}

func Register_user_controllers_all(rg *gin.RouterGroup) {
	router_group := rg.Group("/user")
	router_group.GET("/all", get_all_users)
	router_group.GET("/all/points", get_all_users_with_points)
}
