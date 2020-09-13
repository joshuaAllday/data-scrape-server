package app

import (
	"log"
	"net/http"
	"os"
	"time"

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

	s := http.Server{
		Addr:         port,                                                   // configure the bind address
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(routes), // set the default handler
		ReadTimeout:  5 * time.Second,                                        // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                       // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                      // max time for connections using TCP Keep-Alive
	}

	// start the server
	startServer(&s)
}
