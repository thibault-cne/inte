package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logged() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logged := ctx.MustGet("Logged").(bool)
		if !logged {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
		}
	}
}
