package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/data-scrape/data-scrape-server/api"
	"github.com/data-scrape/data-scrape-server/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupApp() {
	port := os.Getenv("PORT")
	connectionString := os.Getenv("CONNECTION_STRING")

	db, err := db.SetupDatabase(connectionString)

	if err != nil {
		log.Fatalf("Unable to make database connection: %v", err)
	}

	router := mux.NewRouter()
	routes := api.Init(router, db)
	headersOk := handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Authorization", "Accept"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS"})
	fmt.Printf("running on: %s", port)
	http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(routes))

}
