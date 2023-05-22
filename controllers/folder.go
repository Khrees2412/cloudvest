package controllers

import (
	"fmt"

	db "cloudvest/database"
	"cloudvest/models"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func CreateFolder(c *fiber.Ctx) error {
	folder := new(models.Folder)
	userId := fmt.Sprintf("%s", c.Locals("id"))

	if err := c.BodyParser(folder); err != nil {
		log.Error(err)
		return c.JSON(GenericResponse{
			Success: false,
			Message: "Please review your input data",
		})
	}
	folder.UserID = userId

	if err := db.DB.Create(&folder).Error; err != nil {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "Unable to create folder",
			Data:    err.Error(),
		})
	}

	return c.JSON(GenericResponse{
		Success: true,
		Message: fmt.Sprintf("New folder created %s", folder.Name),
	})
}
