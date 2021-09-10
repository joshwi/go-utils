package utils

import (
	"log"
	"os"
	"path/filepath"
)

func Scan(directory string) []string {

	fileList := []string{}

	err := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	return fileList
}
