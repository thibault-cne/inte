package parrainage

import (
	"backend/routes/middlewares"
	parrainage_services "backend/services/parrainage.services"

	"github.com/gin-gonic/gin"
)

func AdminParrainageRoutes(path *gin.RouterGroup) {
	parrainageProcess := parrainage_services.Process

	subpath := path.Group("/parrainage")
	subpath.GET("/toggle", Toggle(parrainageProcess))
	subpath.GET("/pending", middlewares.ParrainageActivated(parrainageProcess), PendingWishes)
	subpath.POST("/grant", middlewares.ParrainageActivated(parrainageProcess), GrantWish)
}
