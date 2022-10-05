package planning

import (
	"github.com/gin-gonic/gin"
)

func PlanningRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/planning")
	subpath.GET("/current", Current)
}
