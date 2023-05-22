package main

import (
	"cloudvest/database"
	"cloudvest/routes"
	"github.com/sirupsen/logrus"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		logrus.Println("No .env file found")
	}
	logrus.Println("Environment variables successfully loaded. Starting application...")
}

func main() {
	app := fiber.New()

	//Connect Database
	database.Connect()

	//Setup routes
	routes.Setup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Cloudvest app set")
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

	if port == "" {
		port = "8000"
	}

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
	// log.Fatal("app started on port:", port)

}
