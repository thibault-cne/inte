package user

import "github.com/gin-gonic/gin"

func AdminUserRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/user")
	subpath.POST("/modify", modify)
}