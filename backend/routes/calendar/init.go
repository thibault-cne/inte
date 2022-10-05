package calendar

import (
	"github.com/gin-gonic/gin"
)

func CalendarRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/calendar")
	subpath.GET("/check/:day", Check)
}
