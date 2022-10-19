package routes

import (
	"backend/routes/admin"
	"backend/routes/auth"
	"backend/routes/calendar"
	"backend/routes/daily"
	"backend/routes/news"
	"backend/routes/planning"
	"backend/routes/stars"
	"backend/routes/suggestions"
	"backend/routes/user"

	"github.com/gin-gonic/gin"
)

func Register(path *gin.RouterGroup) {
	admin.AdminRoutes(path)
	auth.AuthRoutes(path)
	user.UserRoutes(path)
	calendar.CalendarRoutes(path)
	planning.PlanningRoutes(path)
	daily.DailyRoutes(path)
	stars.StarsRoutes(path)
	suggestions.SuggestionsRoutes(path)
	news.NewsRoutes(path)
}
