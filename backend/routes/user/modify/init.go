package modify

import (
	"github.com/gin-gonic/gin"
)

func UserGetRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/modify")
	subpath.POST("/data", Data)
	subpath.POST("/avatar", Avatar)
}
