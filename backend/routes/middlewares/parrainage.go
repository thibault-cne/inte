package middlewares

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParrainageActivated(parrainageProcess *models.ParrainageProcess) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !parrainageProcess.IsProcessOpen {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Parrainage process is closed."})
		}
	}
}
