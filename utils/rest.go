package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

// type Response struct {
// 	Url    string
// 	Type   string
// 	Status string
// 	Data   string
// 	Error  string
// }

func Get(url string) Response {

	errorString := ""

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		errorString = string(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		errorString = string(err.Error())
	}

	output := Response{Url: url, Type: "GET", Status: resp.StatusCode, Data: string(body), Error: errorString}

	return output
}
