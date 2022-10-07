package parrainage

import (
	"backend/models"
	"backend/routes/middlewares"

	"github.com/gin-gonic/gin"
)

func ParrainageRoutes(path *gin.RouterGroup) {
	parrainageProcess := models.Process

	subpath := path.Group("/parrainage", middlewares.NotFirstYear(), middlewares.ParrainageActivated(parrainageProcess))
	subpath.POST("/wish", SetUserWish)
	subpath.GET("/unadopted", GetUnadopted)
	subpath.GET("/endround", EndCurrentRound(parrainageProcess))
}
