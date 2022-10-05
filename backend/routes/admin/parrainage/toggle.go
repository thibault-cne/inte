package parrainage

import (
	parrainageservices "backend/services/parrainage.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Toggle(parrainageProcess *parrainageservices.ParrainageProcess) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		parrainageProcess.IsProcessOpen = !parrainageProcess.IsProcessOpen

		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
