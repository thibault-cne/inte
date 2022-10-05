package parrainage

import (
	parrainage_services "backend/services/parrainage.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUnadopted(ctx *gin.Context) {
	unadoptedUsers := parrainage_services.RetrieveUnadoptedUsers()

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "users": unadoptedUsers})
}
