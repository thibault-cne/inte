package get

import (
	api_services "backend/services/api_response.services"
	users_services "backend/services/users.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserData(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	User, err := users_services.GetUser(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error, user not found"})
	}

	response := api_services.NewProfileDataResponse(User)

	ctx.JSON(http.StatusOK, response)
}
