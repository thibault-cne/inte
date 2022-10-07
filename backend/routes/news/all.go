package news

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// /api/news
func All(ctx *gin.Context) {
	news := models.RetrieveAllNews()

	ctx.JSON(http.StatusOK, news)
}
