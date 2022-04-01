package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/joshwi/go-utils/logger"
)

var write_validation_1 = regexp.MustCompile(`^(\/[a-zA-Z0-9\-\_]{1,20})+$`)
var write_validation_2 = regexp.MustCompile(`^[a-zA-Z0-9\-\_]{0,20}\.(csv|txt|json)$`)
var eof_validation = regexp.MustCompile(`(?i)(\/[a-zA-Z0-9\-\_]+\.\w+$)`)

//Scan a directory for files and subfolders
func Scan(directory string) ([]string, error) {

	output := []string{}

	err := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		rel_path := strings.ReplaceAll(path, directory, "")
		output = append(output, rel_path)
		return err
	})

	if err != nil {
		logger.Logger.Error().Str("directory", directory).Str("status", "Failed").Err(err).Msg("Scan")
		return nil, err
	} else {
		logger.Logger.Info().Str("directory", directory).Str("status", "Success").Msg("Scan")
	}

	return output, nil
}

//Read contents of a file
func Read(filename string) ([]byte, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Logger.Error().Str("file", filename).Str("status", "Failed").Err(err).Msg("Read")
		return nil, err
	} else {
		logger.Logger.Info().Str("file", filename).Str("status", "Success").Msg("Read")
	}

	return data, nil

}

//Write contents to a file
func Write(filename string, data []byte, mode int) error {

	err := ioutil.WriteFile(filename, data, os.FileMode(mode))
	if err != nil {
		logger.Logger.Error().Str("file", filename).Str("status", "Failed").Err(err).Msg("Write")
		return err
	} else {
		logger.Logger.Info().Str("file", filename).Str("status", "Success").Msg("Write")
	}

	return nil

}

//--------------------------------------------------------------------------------------------------------------------------------------------
// Copy directories
//--------------------------------------------------------------------------------------------------------------------------------------------

func Copy(src, dst string) (int64, error) {

	/*
		Input:
			src (string) - Source directory to copy
			dst (string) - Destination directory to copy to
		Output:
			(int64) - Returns number of bytes copied
			(error) - Returns error, otherwise nil
	*/

	response := fmt.Sprintf(`[ Function: Copy ] [ Source: %v ] [ Destination: %v ] [ Status: Success ]`, src, dst)

	filepath := eof_validation.ReplaceAllString(dst, "")

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(filepath, os.FileMode(0766))
		if err != nil {
			response = fmt.Sprintf(`[ Function: Copy ] [ Source: %v ] [ Destination: %v ] [ Status: Failed ] [ Error: %v ]`, src, dst, err)
			log.Fatal(response)
			return 0, err
		}
	}

	_, err = os.Stat(src)
	if err != nil {
		response = fmt.Sprintf(`[ Function: Copy ] [ Source: %v ] [ Destination: %v ] [ Status: Failed ] [ Error: %v ]`, src, dst, err)
		log.Println(response)
		return 0, err
	}

	source, err := os.Open(src)
	if err != nil {
		response = fmt.Sprintf(`[ Function: Copy ] [ Source: %v ] [ Destination: %v ] [ Status: Failed ] [ Error: %v ]`, src, dst, err)
		log.Println(response)
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		response = fmt.Sprintf(`[ Function: Copy ] [ Source: %v ] [ Destination: %v ] [ Status: Failed ] [ Error: %v ]`, src, dst, err)
		log.Println(response)
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	if err != nil {
		response = fmt.Sprintf(`[ Function: Copy ] [ Source: %v ] [ Destination: %v ] [ Status: Failed ] [ Error: %v ]`, src, dst, err)
		log.Println(response)
		return 0, err
	}

	log.Println(response)

	return nBytes, nil
}
