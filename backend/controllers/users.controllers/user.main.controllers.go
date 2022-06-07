package userscontrollers

import "github.com/gin-gonic/gin"

func Register_users_routes(rg *gin.RouterGroup) {
	router_group := rg.Group("/profile")
	Register_user_controllers_all(router_group)
	Register_user_controllers_modify(router_group)
	Register_user_controllers_get(router_group)
	Register_user_controllers_points(router_group)
}
