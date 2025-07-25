package otp

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

type OTPVerifyPayload struct {
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	OTP         string `json:"otp"`
}

func VerifyOTP(countryCode string, phoneNumber string, otp string, clientToken string) bool {
	otpHostUrl := os.Getenv("OTP_HOST_URL")
	otpVerifyApiUrl := os.Getenv("OTP_VERIFY_URL")

	client := resty.New().SetBaseURL(otpHostUrl)

	payload := OTPVerifyPayload{
		CountryCode: countryCode,
		PhoneNumber: phoneNumber,
		OTP:         otp,
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+clientToken).
		SetBody(payload).
		Post(otpVerifyApiUrl)

	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
		return false
	}

	if resp.StatusCode() != http.StatusCreated && resp.StatusCode() != http.StatusOK {
		fmt.Printf("OTP verification failed. Status: %d, Body: %s\n", resp.StatusCode(), resp.String())
		return false
	}

	return true
}
