package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func calcSign(dataToSign string) string {
	// Replace with your actual values
	secret := os.Getenv("CLIENT_SECRET")

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
	clientKey := os.Getenv("CLIENT_ID")
	apiURL := os.Getenv("BASE_URL") + url

	dataToSign := url
	if payload != "" {
		// Compute SHA-256 hash of the payload
		hash := sha256.Sum256([]byte(payload))
		dataToSign += hex.EncodeToString(hash[:])
	}

	signature := calcSign(dataToSign)

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

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// calcSign("/v3/payouts?authorize=false")

	sendRequest("/v3/payouts?authorize=false", `{
		"username": "testuser",
		"accountId": "9db3eb0a-53e4-4bcf-8e3d-8e3a2bdced09",
		"reference": "ref",
		"callbackUrl": "https://psp.stg.01123581.com",
		"requestedAmount": {
			"currency": "XBT",
			"amount": 1
		},
		"destinationAddress": "1BvBMSEYstWetqTFn5Au4m4GFyFei9PT7Z",
		"network": "BITCOIN",
		"verifyBalance": true,
		"feePolicy": "CUSTOM",
		"customFeeRate": 5,
		"subtractFee": true,
		"useCoinConsolidation": true,
		"nonce": 1006
	}`)

	// sendRequest("/v3/payouts?authorize=false", `{
	// 	"username": "testuser",
	// 	"accountId": "f8620c81-d75d-4eff-b2ab-1cb785fbbaaf",
	// 	"reference": "ref",
	// 	"callbackUrl": "https://psp.stg.01123581.com",
	// 	"requestedAmount": {
	// 		"currency": "USDT",
	// 		"amount": 1
	// 	},
	// 	"destinationAddress": "0xA1e8b05C2F2cB4C3C4E3EaF7A7eA6E8A2C8F3B5D",
	// 	"network": "ETHEREUM",
	// 	"verifyBalance": true,
	// 	"feePolicy": "CUSTOM",
	// 	"customFeeRate": 5,
	// 	"subtractFee": true,
	// 	"useCoinConsolidation": true,
	// 	"nonce": 1003
	// }`)
}
