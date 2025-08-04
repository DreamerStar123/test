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
	"time"

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
	// println(signature)

	return signature
}

func sendRequest(method, url, payload string) {
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
	// println(signature)

	// Create HTTP request
	req, err := http.NewRequest(method, apiURL, strings.NewReader(payload))
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

	fmt.Println("Request:", apiURL)
	fmt.Println(payload)

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Response:", string(body))
}

func authorizePayout(payoutId string, otpCode string, currency string) {
	nonce := time.Now().UnixNano()
	username := os.Getenv("USDT_USERNAME")
	if currency == "XBT" {
		username = os.Getenv("XBT_USERNAME")
	}

	url := fmt.Sprintf("/v3/payouts/%s/authorize", payoutId)
	payload := fmt.Sprintf(`{
		"username": "%s",
		"otpCode": "%s",
		"nonce": %v
	}`,
		username,
		otpCode,
		nonce,
	)
	sendRequest("PUT", url, payload)
}

func getPayouts(payoutId string) {
	queryDate := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	url := fmt.Sprintf("/v3/payouts?payoutIds=%s&queryDate=%s", payoutId, queryDate)
	sendRequest("GET", url, "")
}

func createPayout(authorize bool, otpCode string, currency string, amount float64, destinationAddress string, callbackUrl string) {
	nonce := time.Now().UnixNano()

	url := fmt.Sprintf("/v3/payouts?authorize=%v", authorize)
	accountId := accountIdFromCurrency(currency)
	payload := ""

	if currency == "XBT" {
		payload = fmt.Sprintf(`{
			"username": "%s",
			"accountId": "%s",
			"reference": "ref",
			"callbackUrl": "%s",
			"requestedAmount": {
				"currency": "%s",
				"amount": %v
			},
			"destinationAddress": "%s",
			"verifyBalance": true,
			"feePolicy": "SLOW",
			"nonce": %v
		}`, os.Getenv("XBT_USERNAME"), accountId, callbackUrl, currency, amount, destinationAddress, nonce)
	} else {
		payload = fmt.Sprintf(`{
			"username": "%s",
			"otpCode": "%s",
			"accountId": "%s",
			"reference": "ref",
			"callbackUrl": "%s",
			"requestedAmount": {
				"currency": "%s",
				"amount": %v
			},
			"destinationAddress": "%s",
			"verifyBalance": true,
			"feePolicy": "SLOW",
			"nonce": %v
		}`, os.Getenv("USDT_USERNAME"), otpCode, accountId, callbackUrl, currency, amount, destinationAddress, nonce)
	}

	sendRequest("POST", url, payload)
}

func doVoidPayout(payoutId string) {
	nonce := time.Now().UnixNano()

	sendRequest("DELETE", fmt.Sprintf("/v3/payouts/%s?nonce=%d", payoutId, nonce), "")
}

func getDeposits(depositId string) {
	queryDate := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	url := fmt.Sprintf("/v3/deposits?depositIds=%s&queryDate=%s", depositId, queryDate)
	sendRequest("GET", url, "")
}

func accountIdFromCurrency(currency string) string {
	accountId := ""
	switch currency {
	case "XBT":
		accountId = os.Getenv("XBT_ACCOUNT_ID")
	case "USDT":
		accountId = os.Getenv("USDT_ACCOUNT_ID")
	case "USDC":
		accountId = os.Getenv("USDC_ACCOUNT_ID")
	case "LTC":
		accountId = os.Getenv("LTC_ACCOUNT_ID")
	case "BNB":
		accountId = os.Getenv("BNB_ACCOUNT_ID")
	case "ETH":
		accountId = os.Getenv("ETH_ACCOUNT_ID")
	default:
		fmt.Printf("invalid currency %s", currency)
	}
	return accountId
}

func createDeposit(currency string, amount float64, callbackUrl string) {
	nonce := time.Now().UnixNano()

	url := "/v3/deposits"
	payload := fmt.Sprintf(`{
		"accountId": "%s",
		"reference": "ref",
		"callbackUrl": "%s",
		"network": "ETHEREUM",
		"requestedAmount": {
			"currency": "%s",
			"amount": %v
		},
		"nonce": %v
	}`, accountIdFromCurrency(currency), callbackUrl, currency, amount, nonce)

	sendRequest("POST", url, payload)
}

func doVoidDeposit(depositId string) {
	nonce := time.Now().UnixNano()

	url := fmt.Sprintf("/v3/deposits/%s?nonce=%d", depositId, nonce)
	sendRequest("DELETE", url, "")
}

func getAccountsBalance_1() {
	queryDate := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	url := fmt.Sprintf(`/accounts/balance`)
	payload := fmt.Sprintf(`{
		"accountIds": ["%s"],
		"queryDate": "%s"
	}`, os.Getenv("XBT_ACCOUNT_ID"), queryDate)
	sendRequest("POST", url, payload)
}

func getAccountBalance(accountId string) {
	queryDate := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	url := fmt.Sprintf(`/v3/accounts/balance?accountIds=%s&queryDate=%s`, accountId, queryDate)
	fmt.Println(url)
	sendRequest("GET", url, "")
}

func main() {
	const callbackUrl = "https://webhook.site/openpayd"

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// authorizePayout("d9ffbc61-ced2-4b1a-97bb-6d2c4eae4033", "569168", "XBT")
	// getPayouts("120798e3-aef7-4a45-bbee-971be48eb970")
	// createPayout(true, "854773", "XBT", 0.00009073, "bc1qs8juh6zanzmxv99pqpuuln7jl959d6fsj7zf2c", callbackUrl)
	// createPayout(false, "145713", "USDT", 1, "0xBd3a34B02C570BD96bB16950F6b6A868D04747dd", callbackUrl)
	// doVoidPayout("")

	// getDeposits("cd06b124-74f8-42c7-aa8b-6f17ce559b24")
	// createDeposit("XBT", 0.00002, callbackUrl)
	// createDeposit("USDT", 0.01, callbackUrl)
	// createDeposit("ETH", 1, callbackUrl)
	// createDeposit("BNB", 1, callbackUrl)
	// createDeposit("LTC", 1, callbackUrl)
	// createDeposit("USDC", 1, callbackUrl)
	// doVoidDeposit("8aec39ba-9bdf-4492-8c38-b2bf38477754", nonce)

	getAccountBalance(os.Getenv("XBT_ACCOUNT_ID"))
	// getAccountBalance(os.Getenv("USDT_ACCOUNT_ID"))
	// getAccountBalance(os.Getenv("BNB_ACCOUNT_ID"))
}
