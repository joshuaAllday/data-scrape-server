package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) InitHealth() {
	api.BaseRoutes.Router.HandleFunc("/api/health/check/get", api.getHealth).Methods("GET")
}

func (api *API) getHealth(w http.ResponseWriter, r *http.Request) {
	err := api.dbtse.GetHealth()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("pong")

}
