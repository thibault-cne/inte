package stars

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	AmountPerPage = 5
)

func Get(ctx *gin.Context) {
	p := ctx.Param("page")
	page, err := strconv.Atoi(p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number"})
		return
	}

	stars, err := models.GetStarsPage(page*AmountPerPage, AmountPerPage)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	if len(stars) == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"error": "no stars found"})
		return
	}

	ctx.JSON(http.StatusOK, models.NewStarsResponse(stars))
}
