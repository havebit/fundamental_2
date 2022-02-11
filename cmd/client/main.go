package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8081/foobar/15", nil)
	if err != nil {
		log.Panic(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	// resp, err := http.Get("http://localhost:8081/foobar/15")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	var m FoobarResponse
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		log.Panic(err)
	}

	fmt.Println(m)
}

type FoobarResponse struct {
	Foobar string `json:"foobar"`
}
