package get

import (
	"backend/models"
	api_services "backend/services/api_response.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserData(ctx *gin.Context) {
	user := ctx.MustGet("User").(*models.User)

	response := api_services.NewProfileDataResponse(user)

	ctx.JSON(http.StatusOK, response)
}
