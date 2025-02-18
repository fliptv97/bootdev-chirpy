package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}

	if code > 499 {
		log.Printf("Responding with 5xx error: %s\n", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	responseWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")

	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error marshalling JSON: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(code)
	w.Write(data)
}
