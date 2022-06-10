package userscontrollers

import (
	claims_services "backend/services/claims.services"
	users_services "backend/services/users.services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func add_points(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	claims, err := claims_services.RetrieveUserClaims(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	user_id, err := strconv.Atoi(ctx.PostForm("user_id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	points, err := strconv.Atoi(ctx.PostForm("points"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid points"})
		return
	}

	err = users_services.AddPoints(claims.User_id, user_id, points)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func Register_user_controllers_points(rg *gin.RouterGroup) {
	router_group := rg.Group("/points")
	router_group.POST("/add", add_points)
}
