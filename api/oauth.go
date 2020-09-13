package api

import "net/http"

func (api *API) InitOauth() {
	api.BaseRoutes.Router.HandleFunc("/api/auth/token/refresh/get", api.refreshToken).Methods("GET")
}

// Refresh token api

func (api *API) refreshToken(w http.ResponseWriter, r *http.Request) {

}
