package points

import "github.com/gin-gonic/gin"

func UserPointsRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/points")
	subpath.POST("/add", Add)
}
