package suggestions

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	title := ctx.PostForm("title")
	description := ctx.PostForm("description")

	models.NewSuggestions(title, description, userId).Create()

	ctx.JSON(http.StatusOK, gin.H{"message": "Suggestion added"})
}
