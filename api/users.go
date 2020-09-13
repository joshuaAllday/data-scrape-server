package api

import (
	"encoding/json"
	"net/http"

	"github.com/data-scrape/data-scrape-server/models"
)

func (api *API) InitUser() {
	api.BaseRoutes.Router.HandleFunc("/api/user/register/post", api.createUser).Methods("POST")
	api.BaseRoutes.Router.HandleFunc("/api/user/login/post", api.loginUser).Methods("POST")
}

func (api *API) createUser(w http.ResponseWriter, r *http.Request) {
	user := models.UserFromJson(r.Body)
	err := user.SanitizeUserRegister()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	valid, err := api.dbtse.CreateUser(user.Email, models.HashPassword(user.Password))

	if err != nil || *valid == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success")
}

func (api *API) loginUser(w http.ResponseWriter, r *http.Request) {
	user := models.UserFromJson(r.Body)
	err := user.SanitizeUserLogin()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	userDetails, err := api.dbtse.LoginUser(user.Email)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	valid := models.CheckHashPasswords(user.Password, userDetails.Password)

	if !valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Passowrd")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userDetails.ID)
}
