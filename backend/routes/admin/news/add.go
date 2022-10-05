package news

import (
	newsinteservices "backend/services/newsInte.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	content := ctx.PostForm("content")

	newsinteservices.AddNewsInte(newsinteservices.NewNewsInte(content))
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
