package main

import (
	"net/http"
	"fmt"

	"github.com/Brian-Gachiri/rss-agg/internal/database"
	"github.com/Brian-Gachiri/rss-agg/internal/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request)  {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("User not found: %v", err))
			return
		}

		handler(w, r, user)
	}
}