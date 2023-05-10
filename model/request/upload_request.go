package request

import "mime/multipart"

type PathUpload struct {
	FileContent string `json:"file_content" form:"file_content" binding:"required"`
	DirName     string `json:"dir_name" form:"dir_name" binding:"required"`
	NewFilename string `json:"new_filename" form:"new_filename" binding:"required"`
}

type UploadFile struct {
	Files    *multipart.FileHeader `form:"files" binding:"required"`
	FileType string                `form:"file_type,default=default"`
}
