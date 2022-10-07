package logs

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	logs := models.RetrieveAllLogs()

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": logs})
}
