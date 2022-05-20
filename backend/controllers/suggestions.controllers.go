package controllers

import (
	"backend/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func add_suggestion(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	title := ctx.PostForm("title")
	description := ctx.PostForm("description")

	services.RegisterSuggestions(services.NewSuggestions(title, description, claims.User_id))

	ctx.JSON(http.StatusOK, gin.H{"message": "Suggestion added"})
}

func retrieve_all_suggestions(ctx *gin.Context) {
	suggestions, err := services.RetrieveAllSuggestions()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, suggestions)
}

func Register_suggestions_routes(rg *gin.RouterGroup) {
	router_group := rg.Group("/suggestions")
	router_group.POST("/add", add_suggestion)
	router_group.GET("/all", retrieve_all_suggestions)
}
