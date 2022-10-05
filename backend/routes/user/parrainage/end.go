package parrainage

import (
	parrainage_services "backend/services/parrainage.services"

	"github.com/gin-gonic/gin"
)

func EndCurrentRound(parrProcess *parrainage_services.ParrainageProcess) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		parrProcess.IsRoundOpen = false
		parrProcess.CurrentRound += 1

		parrainage_services.EndParrainageRound()
	}
}
