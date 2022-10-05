package user

import (
	"backend/routes/middlewares"
	"backend/routes/user/get"
	"backend/routes/user/modify"
	"backend/routes/user/parrainage"

	"github.com/gin-gonic/gin"
)

func UserRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/user", middlewares.UserStatus(), middlewares.Logged())

	get.UserGetRoutes(subpath)
	modify.UserGetRoutes(subpath)
	parrainage.ParrainageRoutes(subpath)
}
