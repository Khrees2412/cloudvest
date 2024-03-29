package controllers

import (
	"fmt"

	db "cloudvest/database"
	"cloudvest/models"
	"cloudvest/utils"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// kb 204800
var maxByteSize = 209700000 // 200 MB

type GenericResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func StoreFileInFolder(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%s", c.Locals("id"))
	folderId := c.Params("folder")
	if folderId == "" {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "You must specify a folder",
		})
	}

	var folder models.Folder
	if err := db.DB.Where("user_id = ? AND id = ?", userId, folderId).First(&folder).Error; err != nil {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "folder does not exist, you need to create a folder",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Error(err)
		return c.JSON(GenericResponse{
			Success: false,
			Message: "invalid file",
		})
	}
	filesize := file.Size
	filename := file.Filename
	if filesize > int64(maxByteSize) {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "The file size is too large, try something below 200mb",
		})
	}
	data, err := utils.UploadFile(file)

	if err != nil {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "File upload failed",
			Data:    err.Error(),
		})
	}

	newFile := models.File{
		Name:     filename,
		UserID:   userId,
		Url:      data.Location,
		FolderID: folderId,
	}

	if err = db.DB.Create(&newFile).Error; err != nil {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "File upload failed",
		})
	}

	return c.JSON(GenericResponse{
		Success: true,
		Message: fmt.Sprintf("successfully uploaded %s", filename),
		Data:    data.Location,
	})
}
func StoreFile(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%s", c.Locals("id"))

	file, err := c.FormFile("file")
	if err != nil {
		log.Error(err)
		return c.JSON(GenericResponse{
			Success: false,
			Message: "Invalid file",
			Data:    err.Error(),
		})
	}
	filesize := file.Size
	filename := file.Filename
	if filesize > int64(maxByteSize) {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "The file size is too large, try something below 200mb",
		})
	}
	data, err := utils.UploadFile(file)

	if err != nil {
		log.Println(err)
		return c.JSON(GenericResponse{
			Success: false,
			Message: "File upload failed",
		})
	}

	newFile := models.File{
		Name:   filename,
		UserID: userId,
		Url:    data.Location,
	}

	if err = db.DB.Create(&newFile).Error; err != nil {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "File upload failed",
		})
	}

	return c.JSON(GenericResponse{
		Success: true,
		Message: fmt.Sprintf("successfully uploaded %s", filename),
		Data:    data.Location,
	})
}
func GetFiles(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%s", c.Locals("id"))
	var files []models.File

	err := db.DB.Where("id = ?", userId).Find(&files).Error
	if err != nil {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "Unable to find files for user",
		})
	}
	return c.JSON(GenericResponse{
		Success: true,
		Message: "Successfully retrieved files",
		Data:    files,
	})

}

func GetFile(c *fiber.Ctx) error {
	fileName := c.Params("filename")
	userId := fmt.Sprintf("%s", c.Locals("id"))

	var file models.File
	err := db.DB.Where("id = ? AND name = ?", userId, fileName).Find(&file).Error
	if err != nil {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "File not found",
		})
	}
	return c.JSON(GenericResponse{
		Success: true,
		Message: "File successfully retrieved",
		Data:    file,
	})

}

func DeleteFile(c *fiber.Ctx) error {
	fileName := c.Params("filename")
	userId := fmt.Sprintf("%s", c.Locals("id"))

	var file models.File
	err := db.DB.Where("id = ? AND name = ?", userId, fileName).Delete(&file).Error
	if err != nil {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "File couldn't be deleted",
		})
	}
	return c.JSON(GenericResponse{
		Success: true,
		Message: fmt.Sprintf("Deleted %s successfully", fileName),
	})

}

func DownloadFile(c *fiber.Ctx) error {
	fileName := c.Params("filename")
	userId := fmt.Sprintf("%s", c.Locals("id"))
	var file models.File
	err := db.DB.Where("id = ? AND name = ?", userId, fileName).First(&file).Error
	if err != nil {
		return c.JSON(GenericResponse{
			Success: false,
			Message: "File couldn't be downloaded",
		})
	}
	f := file.Url
	return c.Download(f)
}
