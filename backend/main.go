package main

import (
	"backend/controllers"
	"backend/models"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func init_database() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Stars{})
	db.AutoMigrate(&models.DailyGame{})
	db.AutoMigrate(&models.Notifications{})
	db.AutoMigrate(&models.Calendar{})
	db.AutoMigrate(&models.Challenge{})
	db.AutoMigrate(&models.Suggestion{})
	db.AutoMigrate(&models.Tnder{})
}

func main() {
	LoadEnv()
	APP_DOMAIN := os.Getenv("APP_DOMAIN")
	APP_PORT := os.Getenv("APP_PORT")
	init_database()

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
