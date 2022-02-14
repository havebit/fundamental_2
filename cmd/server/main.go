package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pallat/hello/foo"
)

func main() {
	r := mux.NewRouter()

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, world!",
		})
	}

	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/foobar/{param}", foobarHandler).Methods(http.MethodGet)

	log.Println("listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func foobarHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	json.NewEncoder(w).Encode(map[string]string{
		"foobar": foo.SayAny(vars["param"]),
	})
}
