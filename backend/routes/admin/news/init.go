package news

import "github.com/gin-gonic/gin"

func AdminNewsRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/newsInte")

	subpath.POST("/new", Add)
	subpath.POST("/edit", Edit)
	subpath.POST("/delete", Delete)
}
