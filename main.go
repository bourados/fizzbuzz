package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

// DoFizzBuzz this function Returns a list of strings with numbers from 1 to limit,
// where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2,
// all multiples of int1 and int2 are replaced by str1str2
func DoFizzBuzz(w http.ResponseWriter, r *http.Request, history *map[string]int) {
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

	// add current request to history
	curRequest := strconv.Itoa(input.Int1) + "," + strconv.Itoa(input.Int2) + "," + strconv.Itoa(input.Limit) + "," + input.Str1 + "," + input.Str2
	if val, ok := (*history)[curRequest]; ok {
		(*history)[curRequest] = val + 1
	} else {
		(*history)[curRequest] = 1
	}

	w.Write(resp)
}

// DoStatics allows users to know what the most frequent request has been
func DoStatics(w http.ResponseWriter, r *http.Request, history *map[string]int) {
	if len(*history) == 0 {
		w.Write([]byte("There is no query sent yet"))
	} else {
		var mostUsedParams string
		numHits := 0
		// loop over history to get the most used request
		for key, val := range *history {
			if val > numHits {
				mostUsedParams = key
				numHits = val
			}
		}
		w.Write([]byte(mostUsedParams + " : " + strconv.Itoa(numHits) + " times"))
	}
}
