package controllers

import (
	usersservices "backend/services/users.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func retrieveLogs(ctx *gin.Context) {
	logs := usersservices.RetrieveAllLogs()

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": logs})
}

func registerAdminLogsRoutes(rg *gin.RouterGroup) {
	router_group := rg.Group("/logs")

	router_group.GET("/all", retrieveLogs)
}
