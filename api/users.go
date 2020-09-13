package api

import (
	"encoding/json"
	"net/http"

	"github.com/data-scrape/data-scrape-server/models"
)

func (api *API) InitUser() {
	api.BaseRoutes.Router.HandleFunc("/test", api.createUser).Methods("GET")
}

func (api *API) createUser(w http.ResponseWriter, r *http.Request) {
	user, err := models.UserFromJson(r.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	id, err := api.dbtse.CreateUser(user.Email, models.HashPassword(user.Password))

	if err != nil || *id == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success")
}
