package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func calcSign(dataToSign string) string {
	// Replace with your actual values
	secret := "YzNpUzVKLyNMeiZaRWJgdyRcWVNdRn5lOmt8RWptUSZ9YFZYTS1AOHxsTG9GcXh9Ryp+J1R5O28xMX1TZEpBbz9MdWhnZHNZIXMrKFI3LEZocSll"

	// Decode the base64-encoded secret
	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Println("Error decoding secret:", err)
		return "error"
	}

	// Compute HMAC-SHA512 signature
	h := hmac.New(sha512.New, decodedSecret)
	h.Write([]byte(dataToSign))
	signature := hex.EncodeToString(h.Sum(nil))
	println(signature)

	return signature
}

func sendRequest(url, payload string) {
	// Replace with your actual values
	clientKey := "0ee4453c-d8ef-4e39-9deb-a897acc74713"
	apiURL := "https://psp.stg.01123581.com" + url

	// Generate nonce (current timestamp in nanoseconds)
	nonce := fmt.Sprintf("%d", time.Now().UnixNano())

	signature := calcSign(url)

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("key", clientKey)
	req.Header.Set("signature", signature)
	req.Header.Set("nonce", nonce)

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

	fmt.Println("Response:", string(body))
}

func main() {
	// calcSign("/v3/payouts?authorize=false")
	sendRequest("/v3/payouts?authorize=false", `{"username": "testuser"}`)
}
