package main

import (
	"cloudvest/database"
	"cloudvest/logger"
	"cloudvest/routes"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {

	// set up logger
	logger.SetupLogger()
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	name := envs["AWS_BUCKET"]
	editor := envs["AWS_ACCESS_KEY"]

	log.Println("Environment variables ", name, editor)
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
	log.Println("Environment variables successfully loaded. Starting application...")
}

func main() {
	app := fiber.New()

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
	if port == "" {
		port = "8000"
	}

	app.Listen(":" + port)
	// log.Fatal("app started on port:", port)

}
