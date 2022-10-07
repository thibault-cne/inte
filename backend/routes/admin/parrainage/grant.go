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
	godFatherName := ctx.PostForm("godFatherName")
	stepSonName := ctx.PostForm("stepSonName")

	if !models.CheckUserByName(godFatherName) && !models.CheckUserByName(stepSonName) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong users name"})
	}

	models.GrantWishByNames(godFatherName, stepSonName)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
