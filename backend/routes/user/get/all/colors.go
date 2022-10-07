package all

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllWithColors(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.RetrieveAllUsersData())
}
