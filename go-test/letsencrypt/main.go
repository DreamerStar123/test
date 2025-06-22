package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, HTTPS with Let's Encrypt!"))
	})

	// Set up the autocert manager
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("testsslcert.com"), // Use your domain
		Cache:      autocert.DirCache("certs"),                // Directory for storing certificates
	}

	// Create an HTTPS server
	server := &http.Server{
		Addr:      ":443",
		Handler:   m,
		TLSConfig: certManager.TLSConfig(),
	}

	// Start the server
	log.Println("Starting server on :443...")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("ListenAndServeTLS failed: %v", err)
	}
}
