package controllers

import (
	"fmt"
	"os"

	db "risevest/database"
	"risevest/models"
	"risevest/utils"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

var (
	AWS_S3_REGION  = os.Getenv("AWS_REGION")
	AWS_S3_BUCKET  = os.Getenv("AWS_BUCKET")
	AWS_ACCESS_KEY = os.Getenv("AWS_ACCESS_KEY")
	AWS_SECRET_KEY = os.Getenv("AWS_SECRET_KEY")
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

	fmt.Println("aws: ", AWS_ACCESS_KEY)

	user := &models.User{}
	db.DB.Where("uuid = ?", user_id).First(&user)

	folder := &models.Folder{}
	db.DB.Where("user_id = ?", user_id).First(&folder)

	if folder.Name != f {
		return c.JSON(fiber.Map{
			"message": "folder does not exist, you need to create a folder"})
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
	data, uploaderr := utils.UploadFile(file)

	if uploaderr != nil {
		fmt.Println(uploaderr)

		return c.JSON(
			fiber.Map{
				"message": "File upload failed",
				"error":   uploaderr,
			})
	}
	new_file := &models.File{}
	new_file.Name = filename
	new_file.Url = data.Location
	new_file.UserID = user_id

	folder.Files = append(folder.Files, *new_file)

	return c.JSON(fiber.Map{
		"message":  fmt.Sprintf("successfully uploaded %s", filename),
		"file_url": data.Location,
	})
}
func StoreFile(c *fiber.Ctx) error {
	user_id := fmt.Sprintf("%s", c.Locals("id"))
	fmt.Println("aws: ", AWS_ACCESS_KEY)

	user := &models.User{}
	db.DB.Where("uuid = ?", user_id).First(&user)

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
	data, uploaderr := utils.UploadFile(file)

	if uploaderr != nil {
		fmt.Println(uploaderr)
		return c.JSON(fiber.Map{
			"error":   uploaderr,
			"message": "File upload failed",
		})
	}
	new_file := &models.File{}
	new_file.Name = filename
	new_file.Url = data.Location
	new_file.UserID = user_id

	return c.JSON(fiber.Map{
		"message":  fmt.Sprintf("successfully uploaded %s", filename),
		"file_url": data.Location,
	})
}
func GetFiles(c *fiber.Ctx) error {
	user_id := fmt.Sprintf("%s", c.Locals("id"))
	user := &models.User{}

	db.DB.Where("uuid = ?", user_id).Find(&user.Files)
	if len(user.Files) < 1 {
		return c.JSON(fiber.Map{
			"message": "No files found for user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Successfully retrieved files",
		"files":   user.Files,
	})

}

func GetFile(c *fiber.Ctx) error {
	file_name := fmt.Sprintf("%s", c.Locals("filename"))
	return c.JSON(file_name)

}

func DeleteFile(c *fiber.Ctx) error {
	file_name := fmt.Sprintf("%s", c.Locals("filename"))
	return c.JSON(file_name)

}

func DownloadFile(c *fiber.Ctx) error {
	file := c.Params("")
	return c.JSON(file)
}
