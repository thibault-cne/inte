package controllers

import (
	planningservices "backend/services/planning.services"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func addPlanning(ctx *gin.Context) {
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
	err = planningservices.AddPlaning(planningservices.NewPlaning(pictureName, beginingDate, endDate))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error while adding the planning"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func retriveCurrentPlanning(ctx *gin.Context) {
	currentPlanning, err := planningservices.RetrieveLastPlanning()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	// planningResponse, err := api_response_services.NewPlanningResponse(currentPlanning)

	fileContent, err := ioutil.ReadFile("static/images/planning_pictures/" + currentPlanning.Picture)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.Data(http.StatusOK, "image/jpeg", fileContent)
}

func modifyPlanning(ctx *gin.Context) {
	id := ctx.PostForm("id")
	beginingDate := ctx.PostForm("beginingDate")
	endDate := ctx.PostForm("endDate")

	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "You should bind an id to modify the planning"})
	}

	planning := planningservices.RetrievePlanningById(id)

	if planning == nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Wrong id"})
	}

	err := planning.ModifyPlanningDates(beginingDate, endDate)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Internal server error"})
	}

	file, err := ctx.FormFile("picture")

	if err == nil {
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

		err = planning.ModifyPlanningPicture(file)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Internal server error"})
		}
	}

	if err.Error() != "request Content-Type isn't multipart/form-data" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong file type"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func registerPlanningRoutes(rg *gin.RouterGroup) {
	route_group := rg.Group("/planning")
	route_group.GET("/current", retriveCurrentPlanning)
}

func registerAdminPlanningRoutes(rg *gin.RouterGroup) {
	route_group := rg.Group("/planning")
	route_group.POST("/add", addPlanning)
	route_group.POST("/modify", modifyPlanning)
}
