package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-errors/errors"
)

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
	// check if body is empty
	p := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil && !errors.Is(err, io.EOF) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Body must be empty
	if len(p) > 0 {
		http.Error(w, "No parameters accepted", http.StatusBadRequest)
		return
	}

	if len(*history) == 0 {
		w.Write([]byte("There is no query sent yet"))
		return
	}

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
