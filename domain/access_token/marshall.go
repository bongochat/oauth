package access_token

import (
	"encoding/json"
	"time"
)

type TokenResponse struct {
	AccessToken string    `json:"access_token"`
	UserId      int64     `json:"user_id"`
	DeviceId    string    `json:"device_id"`
	IsVerified  bool      `json:"is_verified"`
	DateCreated time.Time `json:"date_created"`
}

func (at *AccessToken) Marshall() interface{} {
	tokenJson, _ := json.Marshal(at)
	var tokenResponse TokenResponse
	json.Unmarshal(tokenJson, &tokenResponse)
	return tokenResponse
}
