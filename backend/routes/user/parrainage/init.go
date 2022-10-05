package parrainage

import (
	"backend/routes/middlewares"
	parrainage_services "backend/services/parrainage.services"

	"github.com/gin-gonic/gin"
)

func ParrainageRoutes(path *gin.RouterGroup) {
	parrainageProcess := parrainage_services.Process

	subpath := path.Group("/parrainage", middlewares.NotFirstYear(), middlewares.ParrainageActivated(parrainageProcess))
	subpath.POST("/wish", SetUserWish)
	subpath.GET("/unadopted", GetUnadopted)
	subpath.GET("/endround", EndCurrentRound(parrainageProcess))
}
