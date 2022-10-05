package suggestions

import (
	suggestion_services "backend/services/suggestion.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	title := ctx.PostForm("title")
	description := ctx.PostForm("description")

	suggestion_services.AddSuggestions(suggestion_services.NewSuggestions(title, description, userId))

	ctx.JSON(http.StatusOK, gin.H{"message": "Suggestion added"})
}
