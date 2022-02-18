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
	private := api.Group("/upload")
	_auth.Post("/login", auth.Login)
	_auth.Post("/register", auth.Register)

	private.Use(utils.SecureAuth())

	private.Post("/folder", controllers.CreateFolder)

	private.Post("/folder/:folder/file", controllers.StoreFileInFolder)

	private.Post("/file", controllers.StoreFile)
	private.Get("/files", controllers.GetFiles)
	private.Get("/file/:fileID", controllers.GetFile)
	private.Delete("/file/:fileID", controllers.DeleteFile)

}
