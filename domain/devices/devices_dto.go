package devices

import "time"

type Devices struct {
	AccessToken string    `json:"access_token"`
	UserId      int64     `json:"user_id"`
	ClientId    int64     `json:"client_id,omitempty"`
	DeviceId    string    `json:"device_id"`
	DeviceType  string    `json:"device_type"`
	DeviceModel string    `json:"device_model"`
	IPAddress   string    `json:"ip_address"`
	IsVerified  bool      `json:"is_verified"`
	DateCreated time.Time `json:"date_created"`
}
