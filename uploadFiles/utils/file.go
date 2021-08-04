package utils

import (
	"os"
	"path/filepath"
	"time"
)

func Mkdir(basePath string) (string, error) {
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(basePath, folderName)

	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return folderPath, nil
}
