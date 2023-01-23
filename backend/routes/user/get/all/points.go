package all

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllWithPoints(ctx *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
