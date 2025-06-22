package main

import (
	"net/http"
)

func main() {
	// Load your certificate and key files
	certFile := "server.crt" // Path to your certificate file
	keyFile := "server.key"  // Path to your private key file

	server := &http.Server{
		Addr: ":443",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, HTTPS with manual cert!"))
		}),
	}

	// Optionally, redirect HTTP to HTTPS
	go func() {
		http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
		}))
	}()

	server.ListenAndServeTLS(certFile, keyFile)
}
