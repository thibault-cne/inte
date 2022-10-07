package get

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserData(ctx *gin.Context) {
	user := ctx.MustGet("User").(*models.User)

	response := models.NewProfileDataResponse(user)

	ctx.JSON(http.StatusOK, response)
}
