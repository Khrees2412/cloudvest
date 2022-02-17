package routes

import (
	"risevest/auth"
	"risevest/controllers"
	"risevest/utils"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Base Api end point
	api := app.Group("/api/v1")
	// Authentication end points
	_auth := api.Group("/auth")
	_auth.Post("/login", auth.Login)
	_auth.Post("/register", auth.Register)

	// api.Use(utils.SecureAuth())

	api.Post("/folder", utils.SecureAuth(), controllers.CreateFolder)

	api.Post("/folder/:folder/file", utils.SecureAuth(), controllers.StoreFileInFolder)

	api.Post("/file", utils.SecureAuth(), controllers.StoreFile)
	api.Get("/files", utils.SecureAuth(), controllers.GetFiles)
	api.Get("/file/:fileID", utils.SecureAuth(), controllers.GetFile)
	api.Delete("/file/:fileID", utils.SecureAuth(), controllers.DeleteFile)

}
