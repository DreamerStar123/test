package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AddressResponse struct {
	FinalBalance int64 `json:"final_balance"`
}

func checkBitcoinAddress(address string) (bool, error) {
	url := fmt.Sprintf("https://blockchain.info/q/addressbalance/%s", address)

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	var balance int64
	if err := json.NewDecoder(resp.Body).Decode(&balance); err != nil {
		return false, err
	}

	return true, nil
}

func checkBitcoin(address string) {
	exists, err := checkBitcoinAddress(address)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if exists {
		fmt.Println("The address exists")
	} else {
		fmt.Println("The address does not exist")
	}
}

func checkEthereum(api_key, addr string) {
	client, err := ethclient.Dial(fmt.Sprintf("https://rpc.ankr.com/eth/%s", api_key))
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(addr) // Replace with your address
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Balance: %s wei\n", balance.String())

	if balance.Cmp(big.NewInt(0)) > 0 {
		fmt.Println("✅ Address has balance — it likely exists.")
	} else {
		fmt.Println("⚠️ Address has zero balance — might be unused.")
	}
}

func main() {
	// checkBitcoin("bc1q5g9f3qhjxlw7m9u0s0y43g755paw274pevk7sx")

	checkEthereum("6860b637cfd221d0f33db6a2c8fac325950f1cf010138f306555f903d80cada7", "0xBd3a34B02C570BD96bB16950F6b6A868D04747dd")
}
