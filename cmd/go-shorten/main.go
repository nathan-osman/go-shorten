package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/go-shorten"
)

func main() {

	// Ensure no more than one argument was supplied
	if len(os.Args) > 2 {
		log.Fatalf("Usage: %s CONFIG", os.Args[0])
	}

	// If no arguments were supplied, write the default configuration
	if len(os.Args) < 2 {
		log.Print("No configuration file specified, creating one...")
		if err := shorten.WriteDefaultConfig("config.json"); err != nil {
			log.Fatal(err)
		} else {
			log.Print("Remember to change the admin password")
			return
		}
	}

	// Attempt to load the file specified as the single argument
	c, err := shorten.LoadConfig(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Configuration loaded")

	// Load the database using the filename in the config file.
	d, err := shorten.LoadDatabase(c.Database)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Database initialized")

	// Initialize the HTTP server
	s, err := shorten.NewServer(c, d)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening for requests on %s...", c.Addr)

	// Wait for SIGINT
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	<-ch

	// Shut down the server
	log.Print("Shutting down...")
	s.Stop()
}
