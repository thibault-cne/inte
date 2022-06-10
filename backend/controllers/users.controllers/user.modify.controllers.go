package userscontrollers

import (
	"backend/models"
	claims_services "backend/services/claims.services"
	users_services "backend/services/users.services"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func modify_profile_data(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	claims, err := claims_services.RetrieveUserClaims(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	temp_user := &models.User{}
	temp_user.ID = claims.User_id

	temp_user.Personal_description = ctx.PostForm("description")
	temp_user.Snapchat_id = ctx.PostForm("snapchat")
	temp_user.Instagram_id = ctx.PostForm("instagram")
	temp_user.Facebook_id = ctx.PostForm("facebook")
	temp_user.Google_id = ctx.PostForm("google")
	temp_user.Hometown = ctx.PostForm("hometown")
	temp_user.Studies = ctx.PostForm("studies")

	err = users_services.ModifyProfileData(temp_user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Function to get the user profile picture from the frontend
// Save the picture in the server and then save the path in the database
func modify_profile_picture(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	claims, err := claims_services.RetrieveUserClaims(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Get the file from the frontend
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	// Get the file extension
	fileExtension := filepath.Ext(file.Filename)

	// Check if the file extension is valid
	if !(fileExtension == ".jpg" || fileExtension == ".jpeg" || fileExtension == ".png") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file extension"})
		return
	}

	// Get the file size
	fileSize := file.Size

	// Check if the file size is valid
	if fileSize > 5000000 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File too big"})
		return
	}

	// Get the file content
	fileContent, err := file.Open()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Create a new file
	newFile, err := os.Create("static/images/profile_pictures/profile_picture_" + strconv.Itoa(claims.User_id) + fileExtension)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Copy the file content to the new file
	_, err = io.Copy(newFile, fileContent)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Close the file
	err = newFile.Close()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Modify the profile picture path in the database
	err = users_services.ModifyProfilePicture(claims.User_id, fileExtension)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func Register_user_controllers_modify(rg *gin.RouterGroup) {
	router_group := rg.Group("/user/modify")
	router_group.POST("/data", modify_profile_data)
	router_group.POST("/picture", modify_profile_picture)
}
