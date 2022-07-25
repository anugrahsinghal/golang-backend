package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	res, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}

type errorBody struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	errorBody := errorBody{err.Error()}
	respondWithJSON(w, code, errorBody)
}
