package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Input is the structure of the body of the request
type Input struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

func main() {
	// history contains the requests and their number of hits
	history := make(map[string]int)

	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/fizzbuzz", func(w http.ResponseWriter, r *http.Request) {
		DoFizzBuzz(w, r, &history)
	}).Methods("POST")

	router.HandleFunc("/stat", func(w http.ResponseWriter, r *http.Request) {
		DoStatics(w, r, &history)
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
