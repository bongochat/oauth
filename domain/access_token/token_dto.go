package access_token

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/utils/resterrors"
	"github.com/golang-jwt/jwt/v5"
)

const (
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken  string    `json:"access_token" bson:"accesstoken"`
	UserId       int64     `json:"user_id" bson:"userid"`
	PhoneNumber  string    `json:"phone_number" bson:"phonenumber"`
	ClientId     string    `json:"client_id,omitempty" bson:"clientid,omitempty"`
	ClientSecret string    `json:"client_secret,omitempty" bson:"clientsecret,omitempty"`
	DeviceId     string    `json:"device_id" bson:"deviceid"`
	DeviceType   string    `json:"device_type" bson:"devicetype"`
	DeviceModel  string    `json:"device_model" bson:"devicemodel"`
	IPAddress    net.IP    `json:"ip_address" bson:"ipaddress"`
	IsVerified   bool      `json:"is_verified" bson:"isverified"`
	IsActive     bool      `json:"is_active" bson:"isactive"`
	CreatedAt    time.Time `json:"created_at" bson:"datecreated"`
	UpdatedAt    time.Time `json:"updated_at" bson:"dateupdated"`
}

type AccessTokenRequest struct {
	GrantType   string `json:"grant_type"`
	Scope       string `json:"scope"`
	DeviceId    string `json:"device_id"`
	DeviceType  string `json:"device_type"`
	DeviceModel string `json:"device_model"`
	IPAddress   net.IP `json:"ip_address"`

	CountryId   int8   `json:"country_id"`
	CountryCode string `json:"country_code"`
	// used for password grant type
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`

	// user for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type RegistrationRequest struct {
	DeviceId    string  `json:"device_id"`
	DeviceType  string  `json:"device_type"`
	DeviceModel string  `json:"device_model"`
	IPAddress   net.IP  `json:"ip_address"`
	AppVersion  string  `json:"app_version"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	CountryId   int8    `json:"country_id"`
	PhoneNumber string  `json:"phone_number"`
	Password    string  `json:"password"`
}

func (at *AccessToken) Validate() resterrors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return resterrors.NewBadRequestError("Invalid access token", "")
	}
	return nil
}

func (rr *RegistrationRequest) ValidateRegistration() resterrors.RestError {
	if rr.CountryId <= 0 {
		return resterrors.NewBadRequestError("Please provide country id", "")
	}
	if rr.PhoneNumber == "" {
		return resterrors.NewBadRequestError("Please provide phone number", "")
	}
	if rr.DeviceId == "" {
		return resterrors.NewBadRequestError("Please provide device ID", "")
	}
	if rr.DeviceType == "" {
		return resterrors.NewBadRequestError("Please provide device type", "")
	}
	if rr.DeviceModel == "" {
		return resterrors.NewBadRequestError("Please provide device model", "")
	}
	if rr.IPAddress == nil {
		return resterrors.NewBadRequestError("Please provide IP address", "")
	}
	return nil
}

func (at *AccessTokenRequest) Validate() resterrors.RestError {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return resterrors.NewBadRequestError("Invalid grant type", "")
	}
	if at.GrantType == grantTypePassword {
		if at.PhoneNumber == "" {
			return resterrors.NewBadRequestError("Please provide phone number", "")
		}
		if at.DeviceId == "" {
			return resterrors.NewBadRequestError("Please provide device ID", "")
		}
		if at.DeviceType == "" {
			return resterrors.NewBadRequestError("Please provide device type", "")
		}
		if at.DeviceModel == "" {
			return resterrors.NewBadRequestError("Please provide device model", "")
		}
		if at.IPAddress == nil {
			return resterrors.NewBadRequestError("Please provide IP address", "")
		}
	}
	return nil
}

func GetNewAccessToken(userId int64, deviceId string) AccessToken {
	return AccessToken{
		UserId:   userId,
		DeviceId: deviceId,
	}
}

func GetNewClientAccessToken(clientId string, clientSecret string) AccessToken {
	return AccessToken{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
}

var secretKey = []byte("secret-key")

func (at *AccessToken) Generate() (string, resterrors.RestError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":   at.UserId,
			"device_id": at.DeviceId,
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		logger.ErrorLog(err)
		return "", resterrors.NewInternalServerError("Token generation failed", "", err)
	}

	return tokenString, nil
}

func (at *AccessToken) GenerateClientToken() (string, resterrors.RestError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"client_id":     at.ClientId,
			"client_secret": at.ClientSecret,
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		logger.ErrorLog(err)
		return "", resterrors.NewInternalServerError("Token generation failed", "", err)
	}

	return tokenString, nil
}

func VerifyTokenString(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		logger.ErrorLog(err)
		return err
	}

	if !token.Valid {
		logger.ErrorMsgLog("Invalid token")
		return fmt.Errorf("invalid token")
	}

	return nil
}
