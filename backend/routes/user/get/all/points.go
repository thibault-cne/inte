package all

import (
	api_services "backend/services/api_response.services"
	users_services "backend/services/users.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllWithPoints(ctx *gin.Context) {
	users, err := users_services.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, api_services.NewAllUsersWithPointsResponse(users))
}
