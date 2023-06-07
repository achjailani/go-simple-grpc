package storagex

import "mime/multipart"

// FileUploadObject is a struct
type FileUploadObject struct {
	File     multipart.File `json:"file"`
	FileName string         `json:"file_name"`
}
