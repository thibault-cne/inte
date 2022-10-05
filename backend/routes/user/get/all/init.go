package all

import "github.com/gin-gonic/gin"

func UserGetAllRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/all")

	// Il y avait deux routes qui pointaient sur la même chose je les ais remis :
	subpath.GET("/", AllUsers)
	//subpath.GET("/", AllWithColors)
	subpath.GET("/points", AllWithPoints)
}
