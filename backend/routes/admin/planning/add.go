package planning

import (
	"backend/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	// Get the begining and end date from the frontend
	beginingDate := ctx.PostForm("beginingDate")
	endDate := ctx.PostForm("endDate")

	// Get the file from the frontend
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
	}

	// Get the file extension
	fileExtension := filepath.Ext(file.Filename)

	// Check if the file extension is valid
	if !(fileExtension == ".jpg" || fileExtension == ".jpeg" || fileExtension == ".png") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid file extension"})
	}

	// Get the file size
	fileSize := file.Size

	// Check if the file size is valid
	if fileSize > 5000000 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "File too big"})
	}

	// Get the file content
	fileContent, err := file.Open()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error while opening the file"})
	}

	// Create a new file
	pictureName := "planning_pictures_" + strings.Replace(beginingDate, "/", "-", -1) + "-" + strings.Replace(endDate, "/", "-", -1) + fileExtension
	newFile, err := os.Create("static/images/planning_pictures/" + pictureName)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cannot create the file"})
	}

	// Copy the file content to the new file
	_, err = io.Copy(newFile, fileContent)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error while copying the file"})
	}

	// Close the file
	err = newFile.Close()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error while closing the file"})
	}

	// Add the planning to the database
	err = models.AddPlaning(models.NewPlaning(pictureName, beginingDate, endDate))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error while adding the planning"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
