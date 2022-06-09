package server

import (
	"backend/config"
	"backend/controllers"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	// Init environment variables
	config.InitEnv()

	APP_DOMAIN := os.Getenv("APP_DOMAIN")
	APP_PORT := os.Getenv("APP_PORT")

	router := gin.Default()
	// Config CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	basepath := router.Group("/api/v1")

	// Add all controllers
	controllers.Register_controllers(basepath)

	router.Run(APP_DOMAIN + ":" + APP_PORT)
}
