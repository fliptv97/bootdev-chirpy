package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

const MAX_CHIRP_LENGTH = 140

var profanities = map[string]struct{}{
	"kerfuffle": {},
	"sharbert":  {},
	"fornax":    {},
}

func handleChirpValidation(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
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
		CleanedBody: censorString(params.Body),
	})
}

func censorString(str string) string {
	words := strings.Split(str, " ")

	for i, word := range words {
		if _, ok := profanities[strings.ToLower(word)]; ok {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}
