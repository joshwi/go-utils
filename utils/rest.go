package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string, headers map[string]string) (Response, error) {

	method := "GET"

	output := Response{Url: url, Method: method}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)

	if err != nil {
		output = Response{Url: url, Method: "GET", Status: 404, Data: "", Error: string(err.Error())}
		log.Println(err)
		return output, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return output, err
	}

	output = Response{Url: url, Method: method, Status: resp.StatusCode, Data: string(body)}

	return output, nil
}
