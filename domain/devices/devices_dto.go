package devices

import (
	"net"
	"time"
)

type Devices struct {
	AccessToken string    `json:"access_token"`
	UserId      int64     `json:"user_id"`
	ClientId    string    `json:"client_id,omitempty"`
	DeviceId    string    `json:"device_id"`
	DeviceType  string    `json:"device_type"`
	DeviceModel string    `json:"device_model"`
	IPAddress   net.IP    `json:"ip_address"`
	IsVerified  bool      `json:"is_verified"`
	DateCreated time.Time `json:"date_created"`
}
