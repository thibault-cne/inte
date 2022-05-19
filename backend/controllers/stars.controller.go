package controllers

import (
	"backend-go/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func get_stars(ctx *gin.Context) {
	stars, err := services.GetAllStars()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, services.NewStarsResponse(stars))
}

func add_stars(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	user_id, err := strconv.Atoi(ctx.PostForm("user_id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	star_type, err := strconv.Atoi(ctx.PostForm("star_type"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid star_type"})
		return
	}

	message := ctx.PostForm("message")

	err = services.AddStars(services.NewStars(claims.User_id, user_id, star_type, message))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func Register_stars_routes(rg *gin.RouterGroup) {
	router_group := rg.Group("/stars")
	router_group.GET("/all", get_stars)
	router_group.POST("/add", add_stars)
}
