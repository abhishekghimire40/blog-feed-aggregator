package main

import "net/http"

func handleReadiness() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responsePayload := StatusOk{
			Status: "ok",
		}
		respondWithJSON(w, 200, responsePayload)
	}
}

func handleError() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responsdWithError(w, 500, "Internal Server Error")
	}
}
