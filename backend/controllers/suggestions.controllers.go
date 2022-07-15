package controllers

import (
	suggestion_services "backend/services/suggestion.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func add_suggestion(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	title := ctx.PostForm("title")
	description := ctx.PostForm("description")

	suggestion_services.AddSuggestions(suggestion_services.NewSuggestions(title, description, userId))

	ctx.JSON(http.StatusOK, gin.H{"message": "Suggestion added"})
}

func retrieve_all_suggestions(ctx *gin.Context) {
	suggestions, err := suggestion_services.RetrieveAllSuggestions()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, suggestions)
}

func registerSuggestionsRoutes(rg *gin.RouterGroup) {
	router_group := rg.Group("/suggestions")
	router_group.POST("/add", add_suggestion)
	router_group.GET("/all", retrieve_all_suggestions)
}

func registerAdminSuggestionRoutes(rg *gin.RouterGroup) {
	router_group := rg.Group("/suggestions")
	router_group.GET("/all", retrieve_all_suggestions)
}
