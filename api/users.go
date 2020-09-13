package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *API) InitUser() {
	api.BaseRoutes.Router.HandleFunc("/test", api.createUser).Methods("GET")
}

func (api *API) createUser(w http.ResponseWriter, r *http.Request) {
	api.dbtse.GetUser()
	fmt.Println("running controllers")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success")
}
