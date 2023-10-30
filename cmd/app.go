package cmd

import (
	"github.com/gorilla/mux"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/db/sql/client"
	"log"
	"net/http"
)

const serverAddress = ":8080"

var dbConfig = client.Config{
	Database: "postgres",
	HostName: "localhost",
	PortNos:  5432,
	UserName: "postgres",
	Password: "password",
	SSLMode:  "disable",
}

func Start() error {
	// Create and initialize a new SQL client
	sqlClient := client.NewSqlClient(dbConfig)
	if err := sqlClient.Init(); err != nil {
		return err
	}

	// Create and initialize the main application
	app := ciam.NewApp()
	if err := app.Init(*sqlClient); err != nil { // pass the SQL client by value
		return err
	}

	// Create a new router
	router := mux.NewRouter()
	// Register app routes (this step will connect your app's routes to the main router)
	appRouter := app.GetRouter()
	router.PathPrefix("/").Handler(appRouter)

	// Start the server
	log.Println("Starting server on", serverAddress)
	return http.ListenAndServe(serverAddress, router)
}
