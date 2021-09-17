package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//Scan a directory for files and subfolders
func Scan(directory string) []string {

	/*
		Input:
			(directory) string - Directory to scan
		Output:
			[]string - List of folders and files in directory
	*/

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

//Read contents of a file
func Read(filename string) map[string]interface{} {

	/*
		Input:
			(filename) string - Path of file to read
		Output:
			map[string]interface{} - JSON structured output
	*/

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	var output map[string]interface{}
	json.Unmarshal(data, &output)

	return output

}
