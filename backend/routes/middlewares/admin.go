package middlewares

import (
	"net/http"

	users_services "backend/services/users.services"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIdInterface, _ := ctx.Get("user_id")
		userId := userIdInterface.(int)

		userIsAdmin, _ := users_services.CheckAdmin(userId)

		if !userIsAdmin {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not administrator."})
		}
	}
}
