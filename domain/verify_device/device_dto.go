package verify_device

import "time"

type VerifyDevice struct {
	AccessToken   string    `json:"access_token"`
	AccountNumber int64     `json:"account_number"`
	ClientId      string    `json:"client_id,omitempty"`
	DeviceId      string    `json:"device_id"`
	IsVerified    bool      `json:"is_verified"`
	DateCreated   time.Time `json:"date_created"`
}
