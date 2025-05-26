package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const coingeckoAPI = "https://api.coingecko.com/api/v3/simple/price"

type PriceResponse struct {
	Price map[string]map[string]float64 `json:"price"`
}

func main() {
	// Example: Get the price of Bitcoin in USD
	currency := "usd"
	ids := "bitcoin"

	url := fmt.Sprintf("%s?ids=%s&vs_currencies=%s", coingeckoAPI, ids, currency)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}

	var priceResponse PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	// Access the price
	bitcoinPrice := priceResponse.Price["bitcoin"][currency]
	fmt.Printf("Current Bitcoin price in %s: %.2f\n", currency, bitcoinPrice)
}
