package planning

import (
	"github.com/gin-gonic/gin"
)

func AdminPlanningRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/planning")

	subpath.POST("/add", Add)
	subpath.POST("/modify", Edit)
}
