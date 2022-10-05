package all

import (
	users_services "backend/services/users.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllWithColors(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users_services.RetrieveAllUsersData())
}
