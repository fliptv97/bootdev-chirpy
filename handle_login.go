package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/fliptv97/bootdev-chirpy/internal/auth"
)

func (cfg *apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	if err := decoder.Decode(&params); err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)

		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)

	if err == sql.ErrNoRows {
		responseWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)

		return
	}

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't get user by email", err)

		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)

	if err != nil {
		responseWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)

		return
	}

	responseWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
