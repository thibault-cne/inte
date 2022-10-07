package parrainage

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUnadopted(ctx *gin.Context) {
	unadoptedUsers := models.RetrieveUnadoptedUsers()

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "users": unadoptedUsers})
}
