package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/handler"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/db/sql/client"
	"github.com/hridayakandel/fine-grain-auth/pkg/router"
)

var (
	serverAddress = ":8080"
	dbConfig      = client.Config{
		Database: "postgres",
		HostName: "localhost",
		PortNos:  5432,
		UserName: "postgres",
		Password: "password",
		SSLMode:  "disable",
	}
)

func Start() error {
	// Initialize the database
	dbClient, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %s", err)
		return err
	}

	// Initialize the store handler
	storeHandler := handler.NewStoreHandler(dbClient)

	// Setup the router
	r := router.SetupRouter(storeHandler)

	// Start the server with the router
	http.Handle("/", r)
	err = http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatalf("Error occurred: %s", err)
		return err
	}

	return nil
}

func initializeDatabase() (*client.SqlClient, error) {
	sqlClient := client.NewSqlClient(dbConfig)
	err := sqlClient.Init(context.Background())
	return sqlClient, err
}
