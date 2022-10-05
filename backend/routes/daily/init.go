package daily

import (
	"github.com/gin-gonic/gin"
)

func DailyRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/daily")

	subpath.GET("/check", Check)
	subpath.GET("/play", Play)
}
