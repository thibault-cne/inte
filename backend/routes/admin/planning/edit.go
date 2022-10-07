package planning

import (
	"backend/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Edit(ctx *gin.Context) {
	id := ctx.PostForm("id")
	beginingDate := ctx.PostForm("beginingDate")
	endDate := ctx.PostForm("endDate")

	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "You should bind an id to modify the planning"})
	}

	planning := models.RetrievePlanningById(id)

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
