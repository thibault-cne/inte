package modify

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Data(ctx *gin.Context) {
	user := ctx.MustGet("User").(*models.User)

	user.Personal_description = ctx.PostForm("description")
	user.Snapchat_id = ctx.PostForm("snapchat")
	user.Instagram_id = ctx.PostForm("instagram")
	user.Facebook_id = ctx.PostForm("facebook")
	user.Google_id = ctx.PostForm("google")
	user.Hometown = ctx.PostForm("hometown")
	user.Studies = ctx.PostForm("studies")

	err := models.ModifyProfileData(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
