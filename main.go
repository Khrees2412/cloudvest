package main

import (
	"risevest/database"
	"risevest/logger"
	"risevest/routes"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	log.Println("Environment variables successfully loaded. Starting application...")
}

func main() {
	app := fiber.New()

	// set up logger
	logger.SetupLogger()

	//Connect Database
	database.Connect()

	//Setup routes
	routes.Setup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("You're home, yaay!!")
	})
	//Activate CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if os.Getenv("PORT") == "" {
		port = "8000"
	}

	app.Listen(":" + port)
	// log.Fatal("app started on port:", port)

}