package routes

import (
	"risevest/auth"
	"risevest/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Base Api end point
	api := app.Group("/api/v1")
	// Authentication end points
	_auth := api.Group("/auth")
	_auth.Post("/login", auth.Login)
	_auth.Post("/register", auth.Register)

	api.Post("/folder/:folder/file", controllers.StoreFileInFolder)

}
