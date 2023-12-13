package access_token

import "time"

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	PhoneNumber string `json:"phone_number"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now())
}
