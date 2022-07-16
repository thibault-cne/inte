package controllers

import (
	parrainageservices "backend/services/parrainage.services"
	usersservices "backend/services/users.services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setUserWish(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	currentParr := parrainageservices.RetrieveCurrentParrainage(userId)

	whishUserName := ctx.PostForm("wishUsername")
	wishOrder := ctx.PostForm("wishOrder")

	wishOrderInt, err := strconv.Atoi(wishOrder)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Wish order must be an integer."})
	}

	if !usersservices.CheckUserByName(whishUserName) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "User don't exists."})
	}

	currentParr.AddWhish(whishUserName, wishOrderInt)

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

func registerParraingeRoutes(rg *gin.RouterGroup) {
	parrainageProcess := &parrainageservices.ParrainageProcess{IsProcessOpen: false, CurrentRound: 1, IsRoundOpen: true}

	routerGroup := rg.Group("/parrainage", setUserStatus(), ensureLoggedIn(), ensureUserNotInYearOne())

	// Set parrainage routes
	routerGroup.POST("/setWish", ensureParrainageProcessIsOn(parrainageProcess), setUserWish)
	routerGroup.GET("/toggleProcess", ensureUserIsAdmin(), toggleParrainageProcess(parrainageProcess))
	routerGroup.GET("/retrieveUsers", ensureParrainageProcessIsOn(parrainageProcess), retrieveAllUnadoptedUsers)
	routerGroup.GET("/endCurrentRound", ensureParrainageProcessIsOn(parrainageProcess), endCurrentParrainageRound(parrainageProcess))
}
