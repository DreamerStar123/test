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

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Response:", string(body))
}

func authorizePayout(payoutId string, nonce int64) {
	sendRequest("PUT", fmt.Sprintf("/v3/payouts/%s/authorize", payoutId), fmt.Sprintf(`{
			"username": "%s",
			"otpCode": "%s"
			"nonce": %v,
		}`,
		os.Getenv("USERNAME"),
		os.Getenv("XBT_PASSWORD"),
		nonce),
	)
}

func getPayouts() {
	queryDate := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	url := fmt.Sprintf("/v3/payouts?queryDate=%s", queryDate)
	sendRequest("GET", url, "")
}

func createPayout(authorize bool, nonce int64, currency string, amount float64, destinationAddress string, network string) {
	url := fmt.Sprintf("/v3/payouts?authorize=%v", authorize)
	payload := ""
	if currency == "XBT" {
		payload = fmt.Sprintf(`{
			"username": "%s",
			"otpCode": "%s",
			"accountId": "%s",
			"reference": "ref",
			"callbackUrl": "https://psp.stg.01123581.com",
			"requestedAmount": {
				"currency": "%s",
				"amount": %v
			},
			"destinationAddress": "%s",
			"network": "%s",
			"verifyBalance": true,
			"feePolicy": "SLOW",
			"subtractFee": true,
			"useCoinConsolidation": true,
			"nonce": %v
		}`, os.Getenv("XBT_USERNAME"), os.Getenv("OTP_CODE"), os.Getenv("XBT_ACCOUNT_ID"), currency, amount, destinationAddress, network, nonce)
	} else {
		payload = fmt.Sprintf(`{
			"username": "%s",
			"accountId": "%s",
			"reference": "ref",
			"callbackUrl": "https://psp.stg.01123581.com",
			"requestedAmount": {
				"currency": "%s",
				"amount": %v
			},
			"destinationAddress": "%s",
			"network": "%s",
			"verifyBalance": true,
			"feePolicy": "SLOW",
			"nonce": %v
		}`, os.Getenv("USDT_USERNAME"), os.Getenv("USDT_ACCOUNT_ID"), currency, amount, destinationAddress, network, nonce)
	}

	println(payload)

	sendRequest("POST", url, payload)
}

func doVoidPayout(payoutId string) {
	sendRequest("DELETE", fmt.Sprintf("/v3/payouts/%s?nonce=%d", payoutId, 1009), "")
}

func getDeposits(depositId string) {
	queryDate := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	url := fmt.Sprintf("/v3/deposits?depositIds=%s&queryDate=%s", depositId, queryDate)
	sendRequest("GET", url, "")
}

func createDeposit(nonce int64, currency string, amount float64) {
	url := "/v3/deposits"
	accountId := ""
	if currency == "XBT" {
		accountId = os.Getenv("XBT_ACCOUNT_ID")
	} else {
		accountId = os.Getenv("USDT_ACCOUNT_ID")
	}
	payload := fmt.Sprintf(`{
		"accountId": "%s",
		"reference": "ref",
		"callbackUrl": "https://psp.stg.01123581.com",
		"requestedAmount": {
			"currency": "%s",
			"amount": %v
		},
		"nonce": %v
	}`, accountId, currency, amount, nonce)

	println(payload)

	sendRequest("POST", url, payload)
}

func doVoidDeposit(depositId string, nonce int64) {
	sendRequest("DELETE", fmt.Sprintf("/v3/deposits/%s?nonce=%d", depositId, nonce), "")
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
	// nonce := int64(1101)

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// authorizePayout(os.Getenv("XBT_PAYOUT_ID"))
	// getPayouts()
	// createPayout(false, nonce, "XBT", 1, "1BvBMSEYstWetqTFn5Au4m4GFyFei9PT7Z", "BITCOIN")
	// createPayout(false, nonce, "USDT", 0.1, "0xB5A2e8C0F3c7D6E9b1C4D5F7A8eF0A6C9F4C2E3D", "ETHEREUM")
	// doVoidPayout("")

	// getDeposits("3a5375a0-a4ee-4064-b3eb-afe7779e5909")
	getDeposits("fe090e04-8ec3-43ef-8e44-a23fd3d27658")
	// createDeposit(nonce, "XBT", 1)
	// createDeposit(nonce, "USDT", 0.1)
	// doVoidDeposit("7fb18232-564c-4846-bc31-210d600f5344", nonce)

	// getAccountsBalance_1()
	// getAccountBalance(os.Getenv("XBT_ACCOUNT_ID"))
	// getAccountBalance(os.Getenv("USDT_ACCOUNT_ID"))
}
