package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type StatusOk struct {
	Status string `json:"status"`
}

type StatusError struct {
	Error string `json:"error"`
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %S", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func responsdWithError(w http.ResponseWriter, status int, msg string) {
	errMsg := StatusError{
		Error: msg,
	}
	respondWithJSON(w, status, errMsg)
}
