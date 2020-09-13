package api

import (
	"encoding/json"
	"net/http"

	"github.com/data-scrape/data-scrape-server/models"
)

func (api *API) isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		valid := models.ValidToken(token)
		if valid {
			endpoint(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized access")
		return
	})
}
