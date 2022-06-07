package userscontrollers

import (
	"backend/services"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func get_user_data(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	User, err := services.GetUserData(claims)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error, user not found"})
		return
	}

	response := services.NewProfileDataResponse(User)

	ctx.JSON(http.StatusOK, response)
}

func get_all_users_notifications(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	users, err := services.RetriveAllUserNotification(claims.User_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func get_users_points_and_stars_stats(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	users, err := services.GetUser(claims.User_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	gold_stars_number, err := services.CountStarsType(claims.User_id, 0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	silver_stars_number, err := services.CountStarsType(claims.User_id, 1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	bronze_stars_number, err := services.CountStarsType(claims.User_id, 2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"points": users.Points, "gold_stars": gold_stars_number, "silver_stars": silver_stars_number, "bronze_stars": bronze_stars_number})
}

// Function to get the user profile picture from the database and send it to the frontend
func provide_user_picture(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Get the file extension
	filePath, err := services.GetProfilePicturePath(claims.User_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Get the file content
	fileContent, err := ioutil.ReadFile("static/images/profile_pictures/" + filePath)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error fo read file"})
		return
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
