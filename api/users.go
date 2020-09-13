package api

import (
	"encoding/json"
	"net/http"

	"github.com/data-scrape/data-scrape-server/models"
	"github.com/gorilla/mux"
)

func (api *API) InitUser() {
	api.BaseRoutes.Router.HandleFunc("/api/user/register/post", api.createUser).Methods("POST")
	api.BaseRoutes.Router.HandleFunc("/api/user/login/post", api.loginUser).Methods("POST")
	api.BaseRoutes.Router.Handle("/api/user/details/get", api.isAuthorized(api.fetchUserData)).Methods("GET")
}

func (api *API) createUser(w http.ResponseWriter, r *http.Request) {
	user := models.UserFromJson(r.Body)
	err := user.SanitizeUserRegister()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	valid, err := api.dbtse.CreateUser(user.Email, models.HashPassword(*user.Password))

	if err != nil || *valid == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
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
		return
	}

	userDetails, err := api.dbtse.LoginUser(user.Email)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	valid := models.CheckHashPasswords(*user.Password, userDetails.Password)

	if !valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Passowrd")
		return
	}

	authToken, err := models.CreateJwt(userDetails.ID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = api.dbtse.AddUserOauthToken(userDetails.ID, authToken.Token)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authToken)
}

func (api *API) fetchUserData(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	err := models.SantizeEmail(email)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	info, err := api.dbtse.FetchUserInfo(email)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(info)

}
