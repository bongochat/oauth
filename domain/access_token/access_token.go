package access_token

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bongochat/utils/resterrors"
	"github.com/golang-jwt/jwt/v5"
)

const (
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string    `json:"access_token"`
	UserId      int64     `json:"user_id"`
	ClientId    int64     `json:"client_id,omitempty"`
	DateCreated time.Time `json:"date_created"`
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// used for password grant type
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`

	// user for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessToken) Validate() resterrors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return resterrors.NewBadRequestError("Invalid access token")
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
		return resterrors.NewBadRequestError("Invalid grant type")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId: userId,
	}
}

var secretKey = []byte("secret-key")

func (at *AccessToken) Generate() (string, resterrors.RestError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": at.UserId,
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", resterrors.NewInternalServerError("Token generation failed", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		log.Println(secretKey, "SECRET")
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	log.Println(token)

	return nil
}
