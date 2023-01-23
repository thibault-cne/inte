package user

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func modify(ctx *gin.Context) {
	var dataUser models.User
	user := ctx.MustGet("User").(*models.User)

	err := ctx.MustBindWith(&dataUser, binding.JSON)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cannot parse user"})
	}

	// Get authorized field
	user.FirstName = dataUser.FirstName
	user.LastName = dataUser.LastName
	user.Email = dataUser.Email
	user.Points = dataUser.Points
	user.UserType = dataUser.UserType

	err = user.Save()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}