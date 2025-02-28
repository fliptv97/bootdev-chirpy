package main

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handleGetChirpById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Couldn't parse chirp id", err)

		return
	}

	chirp, err := cfg.db.GetChirpById(r.Context(), uuid)

	if err == sql.ErrNoRows {
		responseWithError(w, http.StatusNotFound, "Couldn't find any chirps by provided id", err)

		return
	}

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't get chirp by id", err)

		return
	}

	responseWithJSON(w, http.StatusOK, Chirp{
		Id:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserId:    chirp.UserID,
	})
}
