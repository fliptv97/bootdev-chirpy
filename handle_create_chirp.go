package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/fliptv97/bootdev-chirpy/internal/database"
	"github.com/google/uuid"
)

const MAX_CHIRP_LENGTH = 140

var profanities = map[string]struct{}{
	"kerfuffle": {},
	"sharbert":  {},
	"fornax":    {},
}

type Chirp struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserId    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handleCreateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string    `json:"body"`
		UserId uuid.UUID `json:"user_id"`
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

	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   params.Body,
		UserID: params.UserId,
	})

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)

		return
	}

	responseWithJSON(w, http.StatusCreated, Chirp{
		Id:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserId:    chirp.UserID,
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
