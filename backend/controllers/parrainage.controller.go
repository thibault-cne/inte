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

func retrievePendingWishes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "pendingWishes": parrainageservices.RetrievePendingWishes()})
}

// This functions grants wishes for the parrainage process.
// Use a POST request with the following fields :
// `godFatherName` and `stepSonName`
func grantUserWish(ctx *gin.Context) {
	godFatherName := ctx.PostForm("godFatherName")
	stepSonName := ctx.PostForm("stepSonName")

	if !usersservices.CheckUserByName(godFatherName) && !usersservices.CheckUserByName(stepSonName) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong users name"})
	}

	parrainageservices.GrantWishByNames(godFatherName, stepSonName)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func registerParrainageRoutes(rg *gin.RouterGroup) {
	parrainageProcess := &parrainageservices.ParrainageProcess{IsProcessOpen: false, CurrentRound: 1, IsRoundOpen: true}

	usersRouterGroup := rg.Group("/users", setUserStatus(), ensureLoggedIn())
	usersRouterGroup = usersRouterGroup.Group("/parrainage", ensureUserNotInYearOne(), ensureParrainageProcessIsOn(parrainageProcess))

	adminRouterGroup := rg.Group("/admin", setUserStatus(), ensureLoggedIn(), ensureUserIsAdmin())
	adminRouterGroup = adminRouterGroup.Group("/parrainage")

	// Set parrainage routes
	// Users Routes
	usersRouterGroup.POST("/setWish", setUserWish)
	usersRouterGroup.GET("/retrieveUsers", retrieveAllUnadoptedUsers)
	usersRouterGroup.GET("/endCurrentRound", endCurrentParrainageRound(parrainageProcess))

	// Admin Routes
	adminRouterGroup.GET("/toggleProcess", toggleParrainageProcess(parrainageProcess))
	adminRouterGroup.GET("/pendingWishes", ensureParrainageProcessIsOn(parrainageProcess), retrievePendingWishes)
	adminRouterGroup.POST("/grantUserWish", ensureParrainageProcessIsOn(parrainageProcess), grantUserWish)
}
