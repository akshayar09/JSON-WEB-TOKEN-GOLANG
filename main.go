package main

import (
	"log"
	"os"
	"restAPI/controllers"
	"restAPI/database"
	"restAPI/middleware"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("database.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/* database.Connect("root:abc@tcp(127.0.0.1:3306)/rest_api?parseTime=true")
	database.Migrate() */

	// Get database credentials from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Initialize Database
	database.Connect(dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8080")

}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)

		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}

	}
	return router
}
