package parrainage

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PendingWishes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "pendingWishes": models.RetrievePendingWishes()})
}
