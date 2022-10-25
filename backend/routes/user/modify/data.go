package modify

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Data(ctx *gin.Context) {
	user := ctx.MustGet("User").(*models.User)

	user.PersonalDescription = ctx.PostForm("description")
	user.SnapchatId = ctx.PostForm("snapchat")
	user.InstagramId = ctx.PostForm("instagram")
	user.FacebookId = ctx.PostForm("facebook")
	user.GoogleId = ctx.PostForm("google")
	user.Hometown = ctx.PostForm("hometown")
	user.Studies = ctx.PostForm("studies")

	err := models.ModifyProfileData(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
