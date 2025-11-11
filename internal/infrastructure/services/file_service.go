package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FileService struct {
	BasePath string
}

func NewFileService(basePath string) *FileService {
	return &FileService{BasePath: basePath}
}

func (fs *FileService) SaveProfilePic(file multipart.File, filename string) (string, error) {
	uploadDir := filepath.Join(fs.BasePath, "uploads", "profile_pics")
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("cannot create upload directory: %w", err)
	}

	ext := filepath.Ext(filename)
	uniqueName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, uniqueName)

	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot create file: %w", err)
	}
	defer out.Close()

	if _, err = io.Copy(out, file); err != nil {
		return "", fmt.Errorf("cannot write file: %w", err)
	}

	return uniqueName, nil
}
