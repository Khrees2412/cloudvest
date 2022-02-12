package controllers

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func GetFile(c *fiber.Ctx) error {
	log.Info("Hi")
	return c.JSON("")
}

func GetFiles(c *fiber.Ctx) error{
	return c.JSON("")

}

func StoreFile(c *fiber.Ctx) error{
	return c.JSON("")

}

func DeleteFile(c *fiber.Ctx) error{
	return c.JSON("")

}