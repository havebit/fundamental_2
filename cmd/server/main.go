package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/pallat/hello/foo"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, world!",
		})
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/foobar/", foobarHandler)

	log.Println("listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func foobarHandler(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.RequestURI, "/foobar/")

	json.NewEncoder(w).Encode(map[string]string{
		"foobar": foo.SayAny(param),
	})
}
