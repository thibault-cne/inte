package userscontrollers

import (
	api_services "backend/services/api_response.services"
	notifications_services "backend/services/notification.services"
	stars_services "backend/services/stars.services"
	users_services "backend/services/users.services"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func get_user_data(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	User, err := users_services.GetUserData(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error, user not found"})
		return
	}

	response := api_services.NewProfileDataResponse(User)

	ctx.JSON(http.StatusOK, response)
}

func get_all_users_notifications(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	users, err := notifications_services.RetriveAllUserNotification(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func get_users_points_and_stars_stats(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	users, err := users_services.GetUser(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	gold_stars_number, err := stars_services.CountStarsType(userId, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	silver_stars_number, err := stars_services.CountStarsType(userId, 1)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	bronze_stars_number, err := stars_services.CountStarsType(userId, 2)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"points": users.Points, "gold_stars": gold_stars_number, "silver_stars": silver_stars_number, "bronze_stars": bronze_stars_number})
}

// Function to get the user profile picture from the database and send it to the frontend
func provide_user_picture(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("is_logged_in")
	userId := userIdInterface.(int)

	// Get the file extension
	filePath, err := users_services.GetProfilePicturePath(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	// Get the file content
	fileContent, err := ioutil.ReadFile("static/images/profile_pictures/" + filePath)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error fo read file"})
	}

	// Send the file content to the frontend
	ctx.Data(http.StatusOK, "image/jpeg", fileContent)
}

func Register_user_controllers_get(rg *gin.RouterGroup) {
	router_group := rg.Group("/user/get")
	router_group.GET("/data", get_user_data)
	router_group.GET("/notifications", get_all_users_notifications)
	router_group.GET("/stats", get_users_points_and_stars_stats)
	router_group.GET("/picture", provide_user_picture)
}
