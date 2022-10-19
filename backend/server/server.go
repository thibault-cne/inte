package server

import (
	"backend/routes"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	// Init environment variables
	// config.InitEnv()

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
	routes.Register(basepath)

	s := &http.Server{
		Addr:           APP_DOMAIN + ":" + APP_PORT,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start server
	s.ListenAndServe()
}
