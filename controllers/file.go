package controllers

import (
	"fmt"

	db "risevest/database"
	"risevest/models"
	"risevest/utils"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// kb 204800
var maxByteSize = 209700000 // 200 MB

func StoreFileInFolder(c *fiber.Ctx) error {
	user_id := fmt.Sprintf("%s", c.Locals("id"))
	f := c.Params("folder")
	if f == "" {
		return c.JSON(
			fiber.Map{
				"error": "You must specify a folder",
			})
	}
	// user := &models.User{}
	// db.DB.Where("uuid = ?", user_id).First(&user)

	folder := &models.Folder{}
	db.DB.Where("user_id = ?", user_id).First(&folder)

	if folder.Name != f {
		return c.JSON("folder does not exist, you need to create a folder")
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Error(err)
		return c.JSON(
			fiber.Map{
				"error": "Invalid file",
			})
	}
	filesize := file.Size
	filename := file.Filename
	if filesize > int64(maxByteSize) {
		return c.JSON(
			fiber.Map{
				"error": "The file size is too large, try something below 200mb",
			})
	}
	res, uploaderr := utils.UploadFile(file)
	if uploaderr != nil {
		return c.JSON(
			fiber.Map{
				"error": "File upload failed",
			})
	}
	return c.JSON(fiber.Map{
		"message":  fmt.Sprintf("successfully uploaded %s", filename),
		"file_url": res.Url,
	})
}
func StoreFile(c *fiber.Ctx) error {
	return c.JSON(
		fiber.Map{
			"message": " Uploaded successfully",
		})
}
func GetFiles(c *fiber.Ctx) error {
	return c.JSON("")

}

func GetFile(c *fiber.Ctx) error {
	return c.JSON("")

}

func DeleteFile(c *fiber.Ctx) error {
	return c.JSON("")

}
