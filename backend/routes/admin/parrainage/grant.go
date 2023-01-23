package parrainage

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This functions grants wishes for the parrainage process.
// Use a POST request with the following fields :
// `godFatherName` and `stepSonName`
func GrantWish(ctx *gin.Context) {
	godFatherID := ctx.PostForm("godFatherID")
	stepSonID := ctx.PostForm("stepSonID")

	if !models.CheckUserByID(godFatherID) && !models.CheckUserByID(stepSonID) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong users name"})
	}

	models.GrantWish(godFatherID, stepSonID)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
