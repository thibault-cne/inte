package news

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	content := ctx.PostForm("content")

	models.NewNewsInte(content).Create()
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
