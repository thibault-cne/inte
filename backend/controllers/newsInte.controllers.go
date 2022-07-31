package controllers

import (
	newsinteservices "backend/services/newsInte.services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerNewNews(ctx *gin.Context) {
	content := ctx.PostForm("content")

	newsinteservices.AddNewsInte(newsinteservices.NewNewsInte(content))
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func modifyNews(ctx *gin.Context) {
	newsContent := ctx.PostForm("newContent")
	id := ctx.PostForm("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong id. Id must be an integer."})
	}

	newsinteservices.EditNews(intId, newsContent)
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func deleteNews(ctx *gin.Context) {
	id := ctx.PostForm("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong id. Id must be an integer."})
	}

	newsinteservices.DeleteNews(intId)
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func retrieveAllNews(ctx *gin.Context) {
	news := newsinteservices.RetrieveAllNews()

	ctx.JSON(http.StatusOK, news)
}

func registerAdminNewsRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/newsInte")

	routerGroup.POST("/new", registerNewNews)
	routerGroup.POST("/edit", modifyNews)
	routerGroup.POST("/delete", deleteNews)
}

func registerUsersNewsRouter(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/newsInte")

	routerGroup.GET("/all", retrieveAllNews)
}
