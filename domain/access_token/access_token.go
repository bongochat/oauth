package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/bongochat/bongochat-oauth/utils/crypto_utils"
	"github.com/bongochat/bongochat-oauth/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	PhoneNumber string `json:"phone_number"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	PhoneNumber  string `json:"phone_number"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessToken) Validate() *errors.RESTError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("Invalid access token Id")
	}
	return nil
}

func (at *AccessTokenRequest) Validate() *errors.RESTError {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("Invalid grant type")
	}
	return nil
}

func GetNewAccessToken(phoneNumber string) AccessToken {
	return AccessToken{
		PhoneNumber: phoneNumber,
		Expires:     time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMD5(fmt.Sprintf("at-%d-%d-ran", at.PhoneNumber, at.Expires))
}
