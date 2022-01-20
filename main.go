package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var API_URL string = "http://localhost:8080/api/gettopwords"

func main() {

	requestBody, _ := json.Marshal(map[string]string{
		"text": "The HTTP POST method is used to make requests that usually contain a body. It Is used to send data to a server, the data sent is usually used for creating or updating resources.",
	})

	requestBodyBuffer := bytes.NewBuffer(requestBody)

	resp, err := http.Post(API_URL, "application/json", requestBodyBuffer)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
