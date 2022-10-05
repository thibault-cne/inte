package middlewares

import (
	"net/http"

	users_services "backend/services/users.services"

	"github.com/gin-gonic/gin"
)

func NotFirstYear() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIdInterface, _ := ctx.Get("user_id")
		userId := userIdInterface.(int)

		user, _ := users_services.GetUser(userId)

		if user.Current_year == 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to access parrainage process."})
		}
	}
}
