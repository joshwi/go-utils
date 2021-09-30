package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) Response {

	errorString := ""

	resp, err := http.Get(url)

	if err != nil {
		errorString = string(err.Error())
		log.Println(errorString)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		errorString = string(err.Error())
		log.Println(errorString)
	}

	output := Response{Url: url, Type: "GET", Status: resp.StatusCode, Data: string(body), Error: errorString}

	return output
}
