package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	reqBody, err := json.Marshal(map[string]string{
		"name":  "Ivanov Ivan",
		"email": "iivanov@mail.com",
	})

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
