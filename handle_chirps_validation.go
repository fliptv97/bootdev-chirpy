package main

import (
	"encoding/json"
	"net/http"
)

const MAX_CHIRP_LENGTH = 140

func handleChirpValidation(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	if err := decoder.Decode(&params); err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)

		return
	}

	if len(params.Body) > MAX_CHIRP_LENGTH {
		responseWithError(w, http.StatusBadRequest, "Chirp is too long", nil)

		return
	}

	responseWithJSON(w, 200, returnVals{
		Valid: true,
	})
}
