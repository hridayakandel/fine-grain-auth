// cmd/app.go

package cmd

import (
	"context"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/db/sql/client"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/router"
	"log"
	"net/http"
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
	err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %s", err)
		return err
	}

	// Setup routes using the router package
	r := router.SetupRouter()

	http.Handle("/", r)
	err = http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatalf("Error occurred: %s", err)
		return err
	}

	return nil
}

func initializeDatabase() error {
	sqlClient := client.NewSqlClient(dbConfig)
	return sqlClient.Init(context.Background())
}
