package controllers

import (
	users_controllers "backend/controllers/users.controllers"

	"github.com/gin-gonic/gin"
)

func Register_controllers(rg *gin.RouterGroup) {
	// Register users and admin routes
	registerUsersLoggedInRoutes(rg)
	registerAdminRoutes(rg)

	// Register oauth routes
	registerLoginRoutes(rg)
}

func registerUsersLoggedInRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users", setUserStatus(), ensureLoggedIn())

	users_controllers.Register_users_routes(routerGroup)
	registerCalendarsRoutes(routerGroup)
	registerPlanningRoutes(routerGroup)
	registerDailyGameRoutes(routerGroup)
	registerStarsRoutes(routerGroup)
	registerSuggestionsRoutes(routerGroup)
}

func registerAdminRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/admin", ensureUserIsAdmin())

	registerAdminCalendarRoutes(routerGroup)
	registerAdminPlanningRoutes(routerGroup)
	registerAdminStarsRoutes(routerGroup)
	registerAdminSuggestionRoutes(routerGroup)
}
