package modify

import (
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Data(ctx *gin.Context) {
	var dataUser models.User
	user := ctx.MustGet("User").(*models.User)

	err := ctx.MustBindWith(&dataUser, binding.JSON)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cannot parse user"})
	}

	// Get non important fields inside the user
	user.FacebookId = dataUser.FacebookId
	user.InstagramId = dataUser.InstagramId
	user.SnapchatId = dataUser.SnapchatId
	user.GoogleId = dataUser.GoogleId
	user.PersonalDescription = dataUser.PersonalDescription
	user.Hometown = dataUser.Hometown
	user.Studies = dataUser.Studies

	err = user.Save()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	fmt.Println(user.Hometown)
	fmt.Println(user.PersonalDescription)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
