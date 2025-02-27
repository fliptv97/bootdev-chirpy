package main

import "net/http"

func (cfg *apiConfig) handleGetChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetChirps(r.Context())

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)

		return
	}

	formattedChirps := []Chirp{}

	for _, chirp := range chirps {
		formattedChirps = append(formattedChirps, Chirp{
			Id:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			UserId:    chirp.UserID,
			Body:      chirp.Body,
		})
	}

	responseWithJSON(w, http.StatusOK, formattedChirps)
}
