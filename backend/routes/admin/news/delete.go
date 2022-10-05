package news

import (
	newsinteservices "backend/services/newsInte.services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	id := ctx.PostForm("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong id. Id must be an integer."})
	}

	newsinteservices.DeleteNews(intId)
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
