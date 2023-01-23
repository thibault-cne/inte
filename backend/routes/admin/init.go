package admin

import (
	"backend/routes/admin/calendar"
	"backend/routes/admin/logs"
	"backend/routes/admin/news"
	"backend/routes/admin/parrainage"
	"backend/routes/admin/planning"
	"backend/routes/admin/stars"
	"backend/routes/admin/suggestions"
	"backend/routes/admin/user"
	"backend/routes/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/admin", middlewares.UserStatus(), middlewares.Logged(), middlewares.IsAdmin())

	calendar.AdminCalendarRoutes(subpath)
	planning.AdminPlanningRoutes(subpath)
	stars.AdminStarsRoutes(subpath)
	logs.AdminLogsRoutes(subpath)
	news.AdminNewsRoutes(subpath)
	parrainage.AdminParrainageRoutes(subpath)
	suggestions.AdminSuggestionsRoutes(subpath)
	user.AdminUserRoutes(subpath)
}
