package middlewares

import (
	"net/http"

	parrainage_services "backend/services/parrainage.services"

	"github.com/gin-gonic/gin"
)

func ParrainageActivated(parrainageProcess *parrainage_services.ParrainageProcess) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !parrainageProcess.IsProcessOpen {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Parrainage process is closed."})
		}
	}
}
