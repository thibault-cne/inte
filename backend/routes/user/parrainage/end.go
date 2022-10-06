package parrainage

import (
	"backend/models"
	parrainage_services "backend/services/parrainage.services"

	"github.com/gin-gonic/gin"
)

func EndCurrentRound(parrProcess *models.ParrainageProcess) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		parrProcess.IsRoundOpen = false
		parrProcess.CurrentRound += 1

		parrainage_services.EndParrainageRound()
	}
}
