package parrainage

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This functions sets up wishes for the parrainage process.
// Use a POST request with one or all the following fields :
// `firstWish` `secondWish` `thirdWish` `fourthWish` `fifthWish`
// Please notice that the fields correspond to the order of the wishes
func SetUserWish(ctx *gin.Context) {
	user := ctx.MustGet("User").(*models.User)

	currentParr := models.RetrieveCurrentParrainage(user.ID)

	firstWish := ctx.PostForm("firstWish")
	secondWish := ctx.PostForm("secondWish")
	thirdWish := ctx.PostForm("thirdWish")
	fourthWish := ctx.PostForm("fourthWish")
	fifthWish := ctx.PostForm("fifthWish")

	if firstWish != "" && models.CheckUserByName(firstWish) {
		currentParr.AddWhish(firstWish, 1)
	}

	if secondWish != "" && models.CheckUserByName(secondWish) {
		currentParr.AddWhish(secondWish, 2)
	}

	if thirdWish != "" && models.CheckUserByName(thirdWish) {
		currentParr.AddWhish(thirdWish, 3)
	}

	if fourthWish != "" && models.CheckUserByName(fourthWish) {
		currentParr.AddWhish(fourthWish, 4)
	}

	if fifthWish != "" && models.CheckUserByName(fifthWish) {
		currentParr.AddWhish(fifthWish, 5)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
