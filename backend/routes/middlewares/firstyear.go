package middlewares

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func NotFirstYear() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("User").(*models.User)

		if user.Current_year == 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to access parrainage process."})
		}
	}
}
