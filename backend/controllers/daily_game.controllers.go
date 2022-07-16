package controllers

import (
	daily_game_services "backend/services/daily_game.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func check_daily_game(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	status, err := daily_game_services.CheckDailyGame(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": status})
}

func play_daily_game(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	result, err := daily_game_services.PlayDailyGame(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "result": result})
}

func registerDailyGameRoutes(rg *gin.RouterGroup) {
	router_group := rg.Group("/daily-game")
	router_group.GET("/check", check_daily_game)
	router_group.GET("/play", play_daily_game)
}
