package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func sendRequest(method, url, payload string) {
	// Create HTTP request
	apiURL := "https://openexchangerates.org/api" + url
	req, err := http.NewRequest(method, apiURL, strings.NewReader(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Response:", string(body))
}

func getLatest(appId, base string) {
	sendRequest("GET", fmt.Sprintf("/latest.json?app_id=%s&base=%s", appId, base), "")
}

func main() {
	appId := "9b1578c27ec843748dcef280617f5750"

	getLatest(appId, "USD")
}
