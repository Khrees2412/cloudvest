package controllers

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func GetFile(c *fiber.Ctx) error {
	log.Info("Hi")
	// file, err := c.FormFile("file")
	// if err != nil {
	// 	c.JSON(
	// 		fiber.StatusNotAcceptable,
	// 		fiber.Map{
	// 			"error": "Invalid file",
	// 		}
	// 	)
	// }
	// fileInfo, _ := file.Stat()
	// var size = fileInfo.Size()
	return c.JSON("")
}

func GetFiles(c *fiber.Ctx) error {
	return c.JSON("")

}

func StoreFile(c *fiber.Ctx) error {
	return c.JSON("")

}

func DeleteFile(c *fiber.Ctx) error {
	return c.JSON("")

}
