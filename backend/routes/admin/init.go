package admin

import (
	"backend/routes/admin/calendar"
	"backend/routes/admin/planning"
	"backend/routes/admin/stars"
	"backend/routes/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/admin", middlewares.UserStatus(), middlewares.Logged(), middlewares.IsAdmin())

	calendar.AdminCalendarRoutes(subpath)
	planning.AdminPlanningRoutes(subpath)
	stars.AdminStarsRoutes(path)
}
