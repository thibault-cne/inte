package parrainage

import (
	"backend/models"

	"github.com/gin-gonic/gin"
)

func EndCurrentRound(parrProcess *models.ParrainageProcess) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		parrProcess.IsRoundOpen = false
		parrProcess.CurrentRound += 1

		models.EndParrainageRound()
	}
}
