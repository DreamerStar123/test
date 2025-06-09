package main

import (
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// validateURL checks if the provided URL string is valid using regex.
func validateURL(urlStr string) bool {
	re := regexp.MustCompile(`^((https?|ftp):\/\/)?((([a-zA-Z0-9.-]+\.[a-zA-Z]{2,})|(\d{1,3}(\.\d{1,3}){3}))(:\d{1,5})?)(\/[^\s]*)?$`)
	return re.MatchString(urlStr)
}

// Validate Bitcoin Address (Legacy, SegWit, Bech32)
func IsValidBitcoinAddress(address string) bool {
	// Bitcoin Legacy (P2PKH) and SegWit (P2SH) addresses
	btcRegex := regexp.MustCompile(`^(1|3)[a-km-zA-HJ-NP-Z1-9]{25,34}$`)
	if btcRegex.MatchString(address) {
		return true
	}

	// Native SegWit (Bech32) - Starts with "bc1"
	segwitRegex := regexp.MustCompile(`^bc1[0-9a-z]{39,59}$`)
	return segwitRegex.MatchString(address)
}

// Validate Ethereum and Binance Smart Chain (BEP-20) Addresses
func IsValidEthereumBSCAddress(address string) bool {
	// Ethereum and BSC addresses start with "0x" and have 42 characters
	if !strings.HasPrefix(address, "0x") || len(address) != 42 {
		return false
	}
	// Checksum validation (EIP-55 standard)
	return common.IsHexAddress(address)
}

func main() {
	// testURLs := []string{
	// 	"https://www.example.com",
	// 	"http://192.168.1.1",
	// 	"http://example.com:8080/path",
	// 	"https://255.255.255.255",
	// 	"https://www",
	// 	"https://12.23",
	// 	"ftp:/example.com/file",
	// 	"example.com",
	// 	"ftp//example.com/",
	// }

	// for _, u := range testURLs {
	// 	if validateURL(u) {
	// 		// fmt.Printf("'%s' is a valid URL.\n", u)
	// 	} else {
	// 		fmt.Printf("'%s' is not a valid URL.\n", u)
	// 	}
	// }

	valid := IsValidBitcoinAddress("bc1qar0srrr7xfkvy516431ydnw9re59gtzzwf8v9v")
	print(valid)
}
