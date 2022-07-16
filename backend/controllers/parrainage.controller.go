package controllers

import (
	parrainageservices "backend/services/parrainage.services"
	usersservices "backend/services/users.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This functions sets up wishes for the parrainage process.
// Use a POST request with one or all the following fields :
// `firstWish` `secondWish` `thirdWish` `fourthWish` `fifthWish`
// Please notice that the fields correspond to the order of the wishes
func setUserWish(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	currentParr := parrainageservices.RetrieveCurrentParrainage(userId)

	firstWish := ctx.PostForm("firstWish")
	secondWish := ctx.PostForm("secondWish")
	thirdWish := ctx.PostForm("thirdWish")
	fourthWish := ctx.PostForm("fourthWish")
	fifthWish := ctx.PostForm("fifthWish")

	if firstWish != "" && usersservices.CheckUserByName(firstWish) {
		currentParr.AddWhish(firstWish, 1)
	}

	if secondWish != "" && usersservices.CheckUserByName(secondWish) {
		currentParr.AddWhish(secondWish, 2)
	}

	if thirdWish != "" && usersservices.CheckUserByName(thirdWish) {
		currentParr.AddWhish(thirdWish, 3)
	}

	if fourthWish != "" && usersservices.CheckUserByName(fourthWish) {
		currentParr.AddWhish(fourthWish, 4)
	}

	if fifthWish != "" && usersservices.CheckUserByName(fifthWish) {
		currentParr.AddWhish(fifthWish, 5)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func toggleParrainageProcess(parrainageProcess *parrainageservices.ParrainageProcess) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		parrainageProcess.IsProcessOpen = !parrainageProcess.IsProcessOpen

		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

func retrieveAllUnadoptedUsers(ctx *gin.Context) {
	unadoptedUsers := parrainageservices.RetrieveUnadoptedUsers()

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "users": unadoptedUsers})
}

func endCurrentParrainageRound(parrProcess *parrainageservices.ParrainageProcess) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		parrProcess.IsRoundOpen = false
		parrProcess.CurrentRound += 1

		parrainageservices.EndParrainageRound()
	}
}

func registerParrainageRoutes(rg *gin.RouterGroup) {
	parrainageProcess := &parrainageservices.ParrainageProcess{IsProcessOpen: false, CurrentRound: 1, IsRoundOpen: true}

	routerGroup := rg.Group("/parrainage", setUserStatus(), ensureLoggedIn(), ensureUserNotInYearOne())

	// Set parrainage routes
	routerGroup.POST("/setWish", ensureParrainageProcessIsOn(parrainageProcess), setUserWish)
	routerGroup.GET("/toggleProcess", ensureUserIsAdmin(), toggleParrainageProcess(parrainageProcess))
	routerGroup.GET("/retrieveUsers", ensureParrainageProcessIsOn(parrainageProcess), retrieveAllUnadoptedUsers)
	routerGroup.GET("/endCurrentRound", ensureParrainageProcessIsOn(parrainageProcess), endCurrentParrainageRound(parrainageProcess))
}
