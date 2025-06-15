package main

import (
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func getAccessToken() (string, error) {
	token, err := fetchToken(os.Getenv("BASE_URL"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// token, err := getAccessToken()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// println(token)

	auth := AuthTransport{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
	var baseURL *url.URL
	if v := os.Getenv("BASE_URL"); v != "" {
		var err error
		baseURL, err = url.Parse(v)
		if err != nil {
			panic(err)
		}
	}
	NewOpenpayService(auth.Client(), baseURL)
}
