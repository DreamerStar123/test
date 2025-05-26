package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Handle incoming data here
	res := make([]byte, 4096)
	n, _ := r.Body.Read(res)
	res = res[:n]

	println(r.URL.Path, string(res))
	fmt.Fprintf(w, "\nYou've requested: %s\n%s\n", r.URL.Path, res)

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer file.Close()

	// Set the output of the log package to the file
	log.SetOutput(file)

	// Log messages
	log.Println(r.URL.Path, string(res[:n]))
}

func main() {
	http.HandleFunc("/", dataHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
