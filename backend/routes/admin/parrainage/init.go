package parrainage

import (
	"backend/models"
	"backend/routes/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminParrainageRoutes(path *gin.RouterGroup) {
	parrainageProcess := models.Process

	subpath := path.Group("/parrainage")
	subpath.GET("/toggle", Toggle(parrainageProcess))
	subpath.GET("/pending", middlewares.ParrainageActivated(parrainageProcess), PendingWishes)
	subpath.POST("/grant", middlewares.ParrainageActivated(parrainageProcess), GrantWish)
}
