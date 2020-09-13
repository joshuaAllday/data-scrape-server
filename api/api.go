package api

import (
	"github.com/data-scrape/data-scrape-server/store"
	model "github.com/data-scrape/data-scrape-server/store"

	"github.com/gorilla/mux"
)

type Routes struct {
	Router *mux.Router
}

type API struct {
	BaseRoutes *Routes
	dbtse      model.HandlerFunctions
}

func Init(root *mux.Router, db *store.DB) *mux.Router {

	api := &API{
		BaseRoutes: &Routes{},
		dbtse:      db,
	}

	api.BaseRoutes.Router = root

	api.InitUser()
	api.InitHealth()

	return api.BaseRoutes.Router
}
