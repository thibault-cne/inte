package calendar

import (
	"github.com/gin-gonic/gin"
)

func AdminCalendarRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/calendar")

	subpath.POST("/add", AddDay)
	subpath.POST("/modify/:id", Edit)
}
