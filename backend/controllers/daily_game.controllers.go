package controllers

import (
	claims_services "backend/services/claims.services"
	daily_game_services "backend/services/daily_game.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func check_daily_game(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	claims, err := claims_services.RetrieveUserClaims(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	status, err := daily_game_services.CheckDailyGame(claims.User_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": status})
}

func play_daily_game(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	claims, err := claims_services.RetrieveUserClaims(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	result, err := daily_game_services.PlayDailyGame(claims.User_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "result": result})
}

func Register_daily_game_routes(rg *gin.RouterGroup) {
	router_group := rg.Group("/daily-game")
	router_group.GET("/check", check_daily_game)
	router_group.GET("/play", play_daily_game)
}
