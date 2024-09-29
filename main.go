package main

import (
	"fmt"
	"log"
	"net/http"
	
)

func main() {
	// Get environment variables
	//domain := os.Getenv("DOMAIN") // Your domain name (e.g., "yourdomain.com")

	// Define the homepage route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Skill Extraction API!")
	})

	// Start the HTTP server on port 80 (Railway will handle SSL termination)
	httpServer := &http.Server{
		Addr: ":80", // HTTP port 80
		Handler: http.DefaultServeMux,
	}

	log.Fatal(httpServer.ListenAndServe())
}