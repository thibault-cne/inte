package news

import (
	"net/http"
	"strconv"

	newsinte_services "backend/services/newsInte.services"

	"github.com/gin-gonic/gin"
)

func Edit(ctx *gin.Context) {
	newsContent := ctx.PostForm("newContent")
	id := ctx.PostForm("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong id. Id must be an integer."})
	}

	newsinte_services.EditNews(intId, newsContent)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
