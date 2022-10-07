package news

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Edit(ctx *gin.Context) {
	newsContent := ctx.PostForm("newContent")
	id := ctx.PostForm("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong id. Id must be an integer."})
	}

	models.EditNews(intId, newsContent)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
