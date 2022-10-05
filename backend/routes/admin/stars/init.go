package stars

import "github.com/gin-gonic/gin"

func AdminStarsRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/stars")

	subpath.GET("/moderate", Moderate)
}
