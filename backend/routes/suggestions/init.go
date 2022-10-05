package suggestions

import "github.com/gin-gonic/gin"

func SuggestionsRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/suggestions")
	subpath.POST("/", Add)
}
