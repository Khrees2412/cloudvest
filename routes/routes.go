package routes

import (
	"cloudvest/auth"
	"cloudvest/controllers"
	"cloudvest/utils"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Base Api end point
	api := app.Group("/api/v1")
	// Authentication end points
	a := api.Group("/auth")
	private := api.Group("/drive")
	a.Post("/login", auth.Login)
	a.Post("/register", auth.Register)

	private.Use(utils.SecureAuth())

	private.Post("/create-folder", controllers.CreateFolder)

	private.Post("/upload/:folder", controllers.StoreFileInFolder)

	private.Post("/upload", controllers.StoreFile)

	private.Get("/download/:filename", controllers.DownloadFile)

	private.Get("/view/files", controllers.GetFile)

	private.Delete("/:fileID", controllers.DeleteFile)

}
