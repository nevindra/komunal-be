package main

import (
	"log"
	"os"

	"komunal-be/pkg/api"
	db "komunal-be/pkg/api/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger" // Add this import
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize Supabase connection
	if err := db.InitSupabase(); err != nil {
		log.Fatalf("Failed to initialize Supabase: %v", err)
	}

	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Add logger middleware
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} ${latency} ${remote_ip} ${method} ${path}\n",
	}))

	// Setup routes
	api.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
