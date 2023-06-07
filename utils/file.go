package utils

import (
	"mime/multipart"
	"net/http"
)

// GetContentTypeFromFile is a function
func GetContentTypeFromFile(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Reset the file offset back to the beginning
	_, err = file.Seek(0, 0)
	if err != nil {
		return "", err
	}

	// Detect the content type based on the file's magic number
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
