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

func VerifyOTP(countryCode string, phoneNumber string, otp string) bool {
	otpHostUrl := os.Getenv("OTP_HOST_URL")
	otpVerifyApiUrl := os.Getenv("OTP_VERIFY_URL")

	client := resty.New().SetBaseURL(otpHostUrl)

	payload := OTPVerifyPayload{
		CountryCode: countryCode,
		PhoneNumber: phoneNumber,
		OTP:         otp,
	}

	// Replace this with your real token mechanism
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXZpY2VfaWQiOiIiLCJ1c2VyX2lkIjowfQ.Z5orMOpITTFXw_cnKooX2oHSNDEvMrLlPmyCh_fDv-Q"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+token).
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
