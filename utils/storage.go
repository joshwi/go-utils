package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var write_validation_1 = regexp.MustCompile(`^(\/[a-zA-Z0-9\-\_]{1,20})+$`)
var write_validation_2 = regexp.MustCompile(`^[a-zA-Z0-9\-\_]{0,20}\.(csv|txt|json)$`)

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

//Read contents of a file
func ReadTxt(filename string) string {

	/*
		Input:
			(filename) string - Path of file to read
		Output:
			map[string]interface{} - JSON structured output
	*/

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return ``
	}

	return string(data)

}

//Write contents of a file
func Write(filepath string, filename string, data string, mode int) error {

	/*
		Input:
			(filename) string - Path of file to read
		Output:
			map[string]interface{} - JSON structured output
	*/

	response := fmt.Sprintf(`[ Function: Write ] [ Directory: %v ] [ File: %v ] [ Status: Success ]`, filepath, filename)

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		os.MkdirAll(filepath, os.FileMode(mode))
	}

	path := fmt.Sprintf("%v/%v", filepath, filename)

	err = os.WriteFile(path, []byte(data), os.FileMode(mode))

	if err != nil {
		response = fmt.Sprintf(`[ Function: Write ] [ Directory: %v ] [ File: %v ] [ Status: Failed ] [ Error: %v ]`, filepath, filename, err)
		log.Println(response)
		return err
	}

	log.Println(response)

	return nil

}
