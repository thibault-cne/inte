package parrainage

import (
	parrainage_services "backend/services/parrainage.services"
	users_services "backend/services/users.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This functions sets up wishes for the parrainage process.
// Use a POST request with one or all the following fields :
// `firstWish` `secondWish` `thirdWish` `fourthWish` `fifthWish`
// Please notice that the fields correspond to the order of the wishes
func SetUserWish(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	currentParr := parrainage_services.RetrieveCurrentParrainage(userId)

	firstWish := ctx.PostForm("firstWish")
	secondWish := ctx.PostForm("secondWish")
	thirdWish := ctx.PostForm("thirdWish")
	fourthWish := ctx.PostForm("fourthWish")
	fifthWish := ctx.PostForm("fifthWish")

	if firstWish != "" && users_services.CheckUserByName(firstWish) {
		currentParr.AddWhish(firstWish, 1)
	}

	if secondWish != "" && users_services.CheckUserByName(secondWish) {
		currentParr.AddWhish(secondWish, 2)
	}

	if thirdWish != "" && users_services.CheckUserByName(thirdWish) {
		currentParr.AddWhish(thirdWish, 3)
	}

	if fourthWish != "" && users_services.CheckUserByName(fourthWish) {
		currentParr.AddWhish(fourthWish, 4)
	}

	if fifthWish != "" && users_services.CheckUserByName(fifthWish) {
		currentParr.AddWhish(fifthWish, 5)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
