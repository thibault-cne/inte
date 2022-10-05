package stars

import "github.com/gin-gonic/gin"

func StarsRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/stars")

	subpath.GET("/", Get)
	subpath.POST("/add", Add)
}
