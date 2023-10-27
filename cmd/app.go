// app.go in the path: cmd

package cmd

import (
	"context"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/db/sql/client
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	defaultServerAddr = ":8080"
)

var (
	server  *http.Server
	shutSig = make(chan os.Signal, 1)
)

func Start() error {
	// Initialize the database client
	dbConfig := client.Config{
		Database: "postgres",
		HostName: "localhost",
		PortNos:  5432,
		UserName: "postgres",
		Password: "postgres",
		SSLMode:  "disable",
	}
	dbClient := client.NewSqlClient(dbConfig)
	err := dbClient.Init(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := router.SetupRouter()

	server = &http.Server{
		Addr:    defaultServerAddr,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	signal.Notify(shutSig, os.Interrupt)

	<-shutSig // wait for SIGINT (Ctrl+C)
	log.Print("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	return nil
}
