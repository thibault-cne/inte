package middlewares

import (
	claims_services "backend/services/claims.services"

	"github.com/gin-gonic/gin"
)

func UserStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqToken := ctx.Request.Header.Get("Authorization")
		_, err := claims_services.RetrieveUserClaims(reqToken)

		if err != nil {
			ctx.Set("Logged", false)
			return
		}

		ctx.Set("Logged", true)
		// ctx.Set("User", claims.User_id)
	}
}
