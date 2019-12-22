package main

import (
	"log"
	"net/http"
	"encoding/json"
)

func PanicErr(w http.ResponseWriter, err error) {
	if err != nil {
		RespondWithJSON(w, http.StatusServiceUnavailable, nil)
		panic(err)
	}
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Println(err)
		RespondWithJSON(w, http.StatusBadRequest, nil)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}