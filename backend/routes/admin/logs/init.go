package logs

import "github.com/gin-gonic/gin"

func AdminLogsRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/logs")

	subpath.GET("/all", Get)
}
