package models

type Folder struct {
	Base
	Name   string `json:"name"`
	UserID string `json:"user_id"`
	File   []File `json:"files"`
}

type File struct {
	Base
	Name     string `json:"name"`
	UserID   string `json:"user_id"`
	Url      string `json:"url"`
	FolderID string `json:"folder_id" gorm:"foreignKey:folder_id"`
}