package parrainage

import (
	parrainageservices "backend/services/parrainage.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PendingWishes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "pendingWishes": parrainageservices.RetrievePendingWishes()})
}
