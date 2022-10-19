package stars

import (
	"backend/routes/middlewares"

	"github.com/gin-gonic/gin"
)

func StarsRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/stars")

	subpath.GET("/:page", middlewares.UserStatus(), middlewares.Logged(), Get)
	subpath.POST("/add", middlewares.UserStatus(), middlewares.Logged(), middlewares.NotFirstYear(), Add)
}
