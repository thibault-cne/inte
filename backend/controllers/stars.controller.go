package controllers

import (
	api_services "backend/services/api_response.services"
	stars_services "backend/services/stars.services"
	users_services "backend/services/users.services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get_stars(ctx *gin.Context) {
	stars, err := stars_services.GetAllStars()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, api_services.NewStarsResponse(stars))
}

func add_stars(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	// Verify if user is in year > 2
	user, err := users_services.GetUser(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	if user.Current_year == 1 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User must be in year 2 to add stars"})
	}

	user_id, err := strconv.Atoi(ctx.PostForm("user_id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
	}

	star_type, err := strconv.Atoi(ctx.PostForm("star_type"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid star_type"})
	}

	message := ctx.PostForm("message")

	err = stars_services.AddStars(stars_services.NewStars(userId, user_id, star_type, message))

	if err != nil {
		if err.Error() == "message size" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "message size"})
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func moderate_star(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	// Verify if user is admin
	user, err := users_services.GetUser(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	if user.User_type != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User is not admin"})
	}

	star_id, err := strconv.Atoi(ctx.PostForm("star_id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid star_id"})
	}

	err = stars_services.ModerateStar(star_id, userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func registerStarsRoutes(rg *gin.RouterGroup) {
	router_group := rg.Group("/stars")
	router_group.GET("/all", get_stars)
	router_group.POST("/add", add_stars)
	router_group.POST("/moderate", moderate_star)
}

func registerAdminStarsRoutes(rg *gin.RouterGroup) {
	router_group := rg.Group("/stars")
	router_group.POST("/add", add_stars)
	router_group.POST("/moderate", moderate_star)
}
