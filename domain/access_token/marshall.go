package access_token

import (
	"encoding/json"
	"net"
	"time"

	"github.com/bongochat/oauth/users"
)

type TokenResponse struct {
	AccessToken string    `json:"access_token"`
	UserId      int64     `json:"user_id"`
	ClientId    string    `json:"client_id"`
	CountryId   int8      `json:"country_id"`
	CountryCode string    `json:"country_code"`
	PhoneNumber string    `json:"phone_number"`
	DeviceId    string    `json:"device_id"`
	IsVerified  bool      `json:"is_verified"`
	IsActive    bool      `json:"is_active"`
	IPAddress   net.IP    `json:"ip_address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ClientTokenResponse struct {
	AccessToken string    `json:"access_token"`
	ClientId    string    `json:"client_id"`
	IsVerified  bool      `json:"is_verified"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
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
	tokenResponse.PhoneNumber = at.PhoneNumber
	return tokenResponse
}

func (at *AccessToken) ClientTokenMarshall(client *users.Client) interface{} {
	tokenJson, _ := json.Marshal(at)
	clientJson, _ := json.Marshal(client)
	var tokenResponse ClientTokenResponse
	json.Unmarshal(tokenJson, &tokenResponse)
	json.Unmarshal(clientJson, &tokenResponse)
	return tokenResponse
}
