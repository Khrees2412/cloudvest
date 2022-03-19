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

	user_id := fmt.Sprintf("%s", c.Locals("id"))

	if err := c.BodyParser(folder); err != nil {
		log.Error(err)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Please review your input data",
		})
	}
	folder.UserID = user_id

	db.DB.Create(&folder)

	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("New folder created %s", folder.Name),
	})
}
