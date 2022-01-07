package utils

import (
	"archive/zip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

func ReadJson(filename string) ([]byte, error) {
	response := fmt.Sprintf(`[ Function: ReadJson ] [ File: %v ] [ Status: Success ]`, filename)

	// output := [][]string{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		response = fmt.Sprintf(`[ Function: ReadJson ] [ File: %v ] [ Status: Failed ] [ Error: %v ]`, filename, err)
		log.Println(response)
		return []byte{}, err
	}

	log.Println(response)
	return data, nil
}

func WriteTxt(filepath string, filename string, data string, mode int) error {
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

func WriteJson(filepath string, filename string, data []byte, mode int) error {
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

	err = os.WriteFile(path, data, os.FileMode(mode))

	if err != nil {
		response = fmt.Sprintf(`[ Function: Write ] [ Directory: %v ] [ File: %v ] [ Status: Failed ] [ Error: %v ]`, filepath, filename, err)
		log.Println(response)
		return err
	}

	log.Println(response)

	return nil
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

//Read contents of a file
func ReadCsv(filename string) [][]string {

	/*
		Input:
			(filename) string - Path of file to read
		Output:
			map[string]interface{} - JSON structured output
	*/

	response := fmt.Sprintf(`[ Function: ReadCsv ] [ Directory: %v ] [ Status: Success ]`, filename)

	output := [][]string{}

	csv_file, _ := os.Open(filename)
	r := csv.NewReader(csv_file)
	record, err := r.ReadAll()
	if err != nil {
		response = fmt.Sprintf(`[ Function: ReadCsv ] [ Directory: %v ] [ Status: Failed ] [ Error: %v ]`, filename, err)
		log.Println(response)
		return output
	}
	output = record

	return output

}

//Write contents of a file
func WriteCsv(filepath string, filename string, data [][]string, mode int) error {

	/*
		Input:
			(filename) string - Path of file to read
		Output:
			map[string]interface{} - JSON structured output
	*/

	response := fmt.Sprintf(`[ Function: WriteCsv ] [ Directory: %v ] [ File: %v ] [ Status: Success ]`, filepath, filename)

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		os.MkdirAll(filepath, os.FileMode(mode))
	}

	path := fmt.Sprintf("%v/%v", filepath, filename)

	f, err := os.Create(path)
	if err != nil {
		log.Println(err)
	}

	writer := csv.NewWriter(f)

	err = writer.WriteAll(data)

	if err != nil {
		response = fmt.Sprintf(`[ Function: WriteCsv ] [ Directory: %v ] [ File: %v ] [ Status: Failed ] [ Error: %v ]`, filepath, filename, err)
		log.Println(response)
		return err
	}

	log.Println(response)

	return nil

}

func Unzip(src, dest string) error {

	response := fmt.Sprintf(`[ Function: Unzip ] [ Source: %v ] [ Destination: %v ] [ Status: Success ]`, src, dest)

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0777)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		filename := filepath.Base(f.Name)

		path := filepath.Join(dest, filename)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			response = fmt.Sprintf(`[ Function: Unzip ] [ Source: %v ] [ Destination: %v ] [ Status: Failed ] [ Error: %v ]`, src, dest, err)
			log.Println(response)
			return err
		}
	}

	return nil
}

func Zip(src, dest string) error {

	response := fmt.Sprintf(`[ Function: Zip ] [ Source: %v ] [ Destination: %v ] [ Status: Success ]`, src, dest)

	// Get a Buffer to Write To
	outFile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, src, "")

	if err != nil {
		response = fmt.Sprintf(`[ Function: Zip ] [ Source: %v ] [ Destination: %v ] [ Status: Failed ] [ Error: %v ]`, src, dest, err)
		log.Println(response)
		return err
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		response = fmt.Sprintf(`[ Function: Zip ] [ Source: %v ] [ Destination: %v ] [ Status: Failed ] [ Error: %v ]`, src, dest, err)
		log.Println(response)
		return err
	}

	return nil
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		// fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + "/" + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + "/" + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}
