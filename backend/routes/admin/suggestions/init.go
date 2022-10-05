package suggestions

import "github.com/gin-gonic/gin"

func AdminSuggestionsRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/suggestions")
	subpath.GET("/", All)
}
