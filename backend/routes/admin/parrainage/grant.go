package parrainage

import (
	parrainageservices "backend/services/parrainage.services"
	usersservices "backend/services/users.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This functions grants wishes for the parrainage process.
// Use a POST request with the following fields :
// `godFatherName` and `stepSonName`
func GrantWish(ctx *gin.Context) {
	godFatherName := ctx.PostForm("godFatherName")
	stepSonName := ctx.PostForm("stepSonName")

	if !usersservices.CheckUserByName(godFatherName) && !usersservices.CheckUserByName(stepSonName) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong users name"})
	}

	parrainageservices.GrantWishByNames(godFatherName, stepSonName)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
