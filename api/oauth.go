package api

import (
	"encoding/json"
	"net/http"

	"github.com/data-scrape/data-scrape-server/models"
)

func (api *API) InitOauth() {
	api.BaseRoutes.Router.HandleFunc("/api/auth/token/refresh/get", api.refreshToken).Methods("GET")
	api.BaseRoutes.Router.HandleFunc("/api/auth/token/delete/get", api.deleteAuthToken).Methods("GET")

}

func (api *API) refreshToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if len(token) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized access")
		return
	}

	refresh, id, err := models.CreateTokenRefresh(token)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = api.dbtse.AddUserOauthToken(*id, refresh.Token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(refresh)

}

func (api *API) deleteAuthToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if len(token) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized access")
		return
	}
	id, err := models.GetTokenId(token)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = api.dbtse.DeleteAuthToken(*id)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("logged out")

}
