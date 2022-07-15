package userscontrollers

import "github.com/gin-gonic/gin"

func Register_users_routes(rg *gin.RouterGroup) {
	Register_user_controllers_all(rg)
	Register_user_controllers_modify(rg)
	Register_user_controllers_get(rg)
	Register_user_controllers_points(rg)
}
