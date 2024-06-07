package utils

import (
	"io"
	"log"
	"net/http"
)

func GetRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error in making http request: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading the response body: ", err)
	}
	return body
}

func HttpReqest(method, url) []byte {

}