package parrainage

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Toggle(parrainageProcess *models.ParrainageProcess) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		parrainageProcess.IsProcessOpen = !parrainageProcess.IsProcessOpen

		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
