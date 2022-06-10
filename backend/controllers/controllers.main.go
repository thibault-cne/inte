package controllers

import (
	users_controllers "backend/controllers/users.controllers"

	"github.com/gin-gonic/gin"
)

func Register_controllers(rg *gin.RouterGroup) {
	users_controllers.Register_users_routes(rg)
	Register_calendars_routes(rg)
	Register_daily_game_routes(rg)
	Register_login_routes(rg)
	Register_stars_routes(rg)
	Register_suggestions_routes(rg)
}
