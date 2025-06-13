package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"example.com/tree/openpayd/openpayd"
	"github.com/joho/godotenv"
)

func getAccessToken() (string, error) {
	// Replace with your actual values
	clientKey := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	apiURL := os.Getenv("BASE_URL") + "/oauth/token?grant_type=client_credentials"

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	// Set Basic Auth header
	req.SetBasicAuth(clientKey, clientSecret)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	fmt.Println("Request:", apiURL)

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Response:", string(body))

	return "", nil
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// getAccessToken()

	transport := openpayd.AuthTransport{
		ClientKey:    os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
	cl := openpayd.NewClient(transport.Client(), os.Getenv("BASE_URL"))
	req, err := cl.NewRequest("POST", "/oauth/token?grant_type=client_credentials", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cl.Do(req, nil)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	println(string(body))
}
