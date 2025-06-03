package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client := oauth2.NewClient(nil, nil)
		req, err := http.NewRequest("POST", "http://localhost:10001/v1/auth/token", nil)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to reach backend", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
		w.WriteHeader(resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response", http.StatusInternalServerError)
			return
		}
		body = append(body, []byte("  end")...)
		w.Write(body)
	})

	fmt.Println("Server listening on :8081")
	http.ListenAndServe(":8081", nil)
}
