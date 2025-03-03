package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fliptv97/bootdev-chirpy/internal/auth"
	"github.com/fliptv97/bootdev-chirpy/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	if err := decoder.Decode(&params); err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)

		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)

		return
	}

	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		Email:          params.Email,
		HashedPassword: hashedPassword,
	})

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't create user", err)

		return
	}

	responseWithJSON(w, http.StatusCreated, response{
		User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
	})
}
