package news

import (
	newsinteservices "backend/services/newsInte.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// /api/news
func All(ctx *gin.Context) {
	news := newsinteservices.RetrieveAllNews()

	ctx.JSON(http.StatusOK, news)
}
