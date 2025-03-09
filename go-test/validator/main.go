package main

import (
	"fmt"
	"regexp"
)

// validateURL checks if the provided URL string is valid using regex.
func validateURL(urlStr string) bool {
	re := regexp.MustCompile(`^(https?|ftp):\/\/((([a-zA-Z0-9.-]+\.[a-zA-Z]{2,})|(\d{1,3}(\.\d{1,3}){3}))(:\d{1,5})?)(\/[^\s]*)?$`)
	return re.MatchString(urlStr)
}

func main() {
	testURLs := []string{
		"https://www.example.com",
		"http://192.168.1.1",
		"http://example.com:8080/path",
		"https://255.255.255.255",
		"https://www",
		"https://12.23",
		"ftp://example.com/file",
	}

	for _, u := range testURLs {
		if validateURL(u) {
			// fmt.Printf("'%s' is a valid URL.\n", u)
		} else {
			fmt.Printf("'%s' is not a valid URL.\n", u)
		}
	}
}
