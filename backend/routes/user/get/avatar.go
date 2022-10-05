package get

import (
	users_services "backend/services/users.services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Function to get the user profile picture from the database and send it to the frontend
func Avatar(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("is_logged_in")
	userId := userIdInterface.(int)

	// Get the file extension
	filePath, err := users_services.GetProfilePicturePath(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	// Get the file content
	fileContent, err := os.ReadFile("static/images/profile_pictures/" + filePath)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error fo read file"})
	}

	// Send the file content to the frontend
	ctx.Data(http.StatusOK, "image/jpeg", fileContent)
}
