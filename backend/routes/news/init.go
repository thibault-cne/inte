package news

import "github.com/gin-gonic/gin"

func NewsRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/news")

	subpath.GET("/", All)
}
