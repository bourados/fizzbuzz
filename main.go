package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Input struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/fizzbuzz", DoFizzBuzz).Methods("POST")
	router.HandleFunc("/stat", DoStatics)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// DoFizzBuzz this function Returns a list of strings with numbers from 1 to limit,
// where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2,
// all multiples of int1 and int2 are replaced by str1str2
func DoFizzBuzz(w http.ResponseWriter, r *http.Request) {
	var input Input
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&input)

	response := make([]string, input.Limit)

	for i := 1; i < input.Limit+1; i++ {
		r := ""
		if i%(input.Int1*input.Int2) == 0 {
			r = input.Str1 + input.Str2
		} else if i%input.Int2 == 0 {
			r = input.Str2
		} else if i%input.Int1 == 0 {
			r = input.Str1
		} else {
			r = strconv.Itoa(i)
		}
		response[i-1] = r
	}

	resp, err := json.Marshal(response)

	if err != nil {
		log.Println(err)
	}

	w.Write(resp)
}

// Dostatics allows users to know what the most frequent request has been
func DoStatics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Statics to be implemented")
}
