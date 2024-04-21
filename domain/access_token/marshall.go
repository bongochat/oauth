package access_token

import (
	"encoding/json"
	"time"

	"github.com/bongochat/oauth/users"
)

type TokenResponse struct {
	AccessToken string    `json:"access_token"`
	UserId      int64     `json:"user_id"`
	CountryCode string    `json:"country_code"`
	PhoneNumber string    `json:"phone_number"`
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

func (at *AccessToken) TokenMarshall(user *users.User) interface{} {
	tokenJson, _ := json.Marshal(at)
	userJson, _ := json.Marshal(user)
	var tokenResponse TokenResponse
	json.Unmarshal(tokenJson, &tokenResponse)
	json.Unmarshal(userJson, &tokenResponse)
	return tokenResponse
}
