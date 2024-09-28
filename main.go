package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	// Get environment variables
	domain := os.Getenv("DOMAIN")           // Your domain name (e.g., "yourdomain.com")
	certCacheDir := os.Getenv("CERT_CACHE") // Path for certificate cache

	// Fallback if CERT_CACHE is not set
	if certCacheDir == "" {
		certCacheDir = "/certs"
	}

	// Define the homepage route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Skill Extraction API!")
	})

	

	// Configure Let's Encrypt autocert manager
	certManager := autocert.Manager{
		Cache:      autocert.DirCache(certCacheDir), // Directory for storing certificates
		Prompt:     autocert.AcceptTOS,              // Automatically accept Let's Encrypt's Terms of Service
		HostPolicy: autocert.HostWhitelist(domain),  // Replace with your actual domain
	}

	// HTTP server for redirecting HTTP traffic and handling Let's Encrypt HTTP-01 challenges
	go func() {
		httpServer := &http.Server{
			Addr: ":80", // HTTP port 80
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Check if this is a request for an HTTP-01 challenge
				if r.URL.Path == "/.well-known/acme-challenge/" {
					certManager.HTTPHandler(nil).ServeHTTP(w, r)
				} else {
					// Otherwise, redirect to HTTPS
					http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
				}
			}),
		}
		log.Fatal(httpServer.ListenAndServe())
	}()

	// Start the HTTPS server
	httpsServer := &http.Server{
		Addr:      ":443",               // HTTPS port
		Handler:   nil,                  // Your handlers are set up with `http.HandleFunc` above
		TLSConfig: certManager.TLSConfig(), // TLS configuration for Let's Encrypt
	}

	// Start the HTTPS server
	log.Fatal(httpsServer.ListenAndServeTLS("", ""))
}
