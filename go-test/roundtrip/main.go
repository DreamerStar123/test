package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// CustomTransport implements the http.RoundTripper interface
type CustomTransport struct {
	Transport http.RoundTripper // The underlying transport (default or custom)
}

// RoundTrip executes a single HTTP transaction
func (c *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add a custom header

	// Use the underlying transport to perform the request
	return c.Transport.RoundTrip(req)
}

func main() {
	// Create a new HTTP client with the custom transport
	client := &http.Client{
		Transport: &CustomTransport{
			Transport: http.DefaultTransport, // Use the default transport
		},
	}

	// Create a new request
	req, err := http.NewRequest("GET", "https://bing.com", nil)
	if err != nil {
		panic(err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Println("Response Status:", resp.Status)

	// Print the response body (for demonstration)
	var responseBody bytes.Buffer
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body:", responseBody.String())
}
