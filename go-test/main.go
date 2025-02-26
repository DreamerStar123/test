package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Validate Bitcoin Address (Legacy, SegWit, Bech32)
func isValidBitcoinAddress(address string) bool {
	// Bitcoin Legacy (P2PKH) and SegWit (P2SH) addresses
	btcRegex := regexp.MustCompile(`^(1|3)[a-km-zA-HJ-NP-Z1-9]{25,34}$`)
	if btcRegex.MatchString(address) {
		return true
	}

	// Native SegWit (Bech32) - Starts with "bc1"
	segwitRegex := regexp.MustCompile(`^bc1[a-z0-9]{8,87}$`)
	if segwitRegex.MatchString(address) {
		return true
	}

	return false
}

// Validate Ethereum and Binance Smart Chain (BEP-20) Addresses
func isValidEthereumBSCAddress(address string) bool {
	// Ethereum and BSC addresses start with "0x" and have 42 characters
	if !strings.HasPrefix(address, "0x") || len(address) != 42 {
		return false
	}
	// Checksum validation (EIP-55 standard)
	return common.IsHexAddress(address)
}

// Detect Wallet Address Type
func detectAddressType(address string) string {
	if isValidBitcoinAddress(address) {
		return "Bitcoin (BTC)"
	} else if isValidEthereumBSCAddress(address) {
		return "Ethereum (ERC-20) or Binance Smart Chain (BEP-20)"
	}
	return "Unknown or Invalid Address"
}

func main() {
	// Test Addresses
	addresses := []string{
		"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",         // Bitcoin (Legacy)
		"bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwfq3mg", // Bitcoin (Bech32)
		"0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae", // Ethereum (ERC-20)
		"0x8fF795a6F712b267b4Fe4931AFe132E21e406caA", // Binance Smart Chain (BEP-20)
		"abcd1234", // Invalid Address
	}

	if isValidBitcoinAddress("0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae") {
		fmt.Print("valid")
	}

	for _, addr := range addresses {
		fmt.Printf("Address: %s -> %s\n", addr, detectAddressType(addr))
	}
}
