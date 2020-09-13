package api

import (
	"github.com/data-scrape/data-scrape-server/model"

	"github.com/gorilla/mux"
)

type Routes struct {
	Router *mux.Router
}

type API struct {
	BaseRoutes *Routes
	dbtse      model.HandlerFunctions
}

func Init(root *mux.Router, db *model.DB) *mux.Router {

	api := &API{
		BaseRoutes: &Routes{},
		dbtse:      db,
	}

	api.BaseRoutes.Router = root

	api.InitUser()

	return api.BaseRoutes.Router
}
