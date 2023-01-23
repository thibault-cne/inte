package get

import (
	"backend/routes/user/get/all"

	"github.com/gin-gonic/gin"
)

func UserGetRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/get")
	all.UserGetAllRoutes(subpath)

	subpath.GET("/notifs", Notifs)
	subpath.GET("/stats", Stats)
	subpath.GET("/avatar", Avatar)
}
