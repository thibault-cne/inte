package logs

import (
	usersservices "backend/services/users.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	logs := usersservices.RetrieveAllLogs()

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": logs})
}
