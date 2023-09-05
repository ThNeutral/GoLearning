package main

import (
	"fmt"
	"net/http"

	"github.com/ThNeutral/GoLearning/internal/auth"
	"github.com/ThNeutral/GoLearning/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn`t verify apiKey: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Wrong apiKey: %v", err))
			return
		}

		handler(w, r, user)
	}
}
