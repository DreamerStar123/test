package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func main() {
	http.HandleFunc("/generate", generateHandler)
	http.HandleFunc("/validate", validateHandler)
	http.HandleFunc("/qrcode", qrCodeHandler)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	// Generate a new TOTP key
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MyApp",
		AccountName: "user@example.com",
	})
	if err != nil {
		http.Error(w, "Failed to generate key", http.StatusInternalServerError)
		return
	}

	// Display the key and QR code URL
	fmt.Fprintf(w, "Your OTP secret key: %s\n", key.Secret())
	fmt.Fprintf(w, "Scan this QR code with Google Authenticator: %s\n", key.URL())
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	// Get the OTP and secret from the request
	otp := r.URL.Query().Get("otp")
	secret := r.URL.Query().Get("secret")

	// Validate the OTP
	valid := totp.Validate(otp, secret)
	if valid {
		fmt.Fprintln(w, "OTP is valid!")
	} else {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
	}
}

func generateQRCode(secret string) ([]byte, error) {
	// Format the otpauth URL
	otpauth := fmt.Sprintf("otpauth://totp/AccountName?secret=%s&issuer=YourIssuer", secret)

	// Generate QR code
	return qrcode.Encode(otpauth, qrcode.Medium, 256)
}

func qrCodeHandler(w http.ResponseWriter, r *http.Request) {
	secret := r.URL.Query().Get("secret")
	if secret == "" {
		http.Error(w, "Missing secret", http.StatusBadRequest)
		return
	}

	qrCode, err := generateQRCode(secret)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	// Set headers
	w.Header().Set("Content-Type", "image/png")
	w.Write(qrCode)
}
