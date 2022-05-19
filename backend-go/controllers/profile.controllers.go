package controllers

import (
	"backend-go/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func get_user_data(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	User, err := services.GetUserData(claims)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error, user not found"})
		return
	}

	response := services.NewProfileDataResponse(User)

	ctx.JSON(http.StatusOK, response)
}

func Register_profile_routes(rg *gin.RouterGroup) {
	router_group := rg.Group("/profile")
	router_group.GET("/all", get_user_data)
}
