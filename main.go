package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Initialize the database
	db, err := NewDatabase(*dbFilename)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Database initialized")

	// Initialize the HTTP server
	srv, err := NewServer(*addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server listening on %s", *addr)

	// Wait for SIGINT
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c

	// Stop the server
	log.Print("Shutting down...")
	srv.Stop()
	log.Print("Server stopped")
}
