package middlewares

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("User").(*models.User)

		if user.User_type != "admin" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not administrator."})
		}
	}
}
