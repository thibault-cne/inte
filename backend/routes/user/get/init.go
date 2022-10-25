package get

import (
	"github.com/gin-gonic/gin"
)

func UserGetRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/get")
	subpath.GET("/notifs", Notifs)
	subpath.GET("/stats", Stats)
	subpath.GET("/avatar", Avatar)
}
