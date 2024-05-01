package users

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/bongochat/utils/resterrors"
	"github.com/go-resty/resty/v2"
)

type User struct {
	Id          int64  `json:"id"`
	CountryId   int8   `json:"country_id"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

var (
	userHostUrl     = os.Getenv("USER_HOST_URL")
	userLoginApiUrl = os.Getenv("USER_LOGIN_API_URL")
	userClient      = resty.New().SetBaseURL(userHostUrl)
)

func LoginUser(phone_number string, password string) (*User, resterrors.RestError) {
	request := UserLoginRequest{
		PhoneNumber: phone_number,
		Password:    password,
	}

	response, err := userClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(userLoginApiUrl)

	if err != nil {
		return nil, resterrors.NewInternalServerError("Could not login", "", err)
	}

	if response == nil {
		return nil, resterrors.NewInternalServerError("Invalid client response", "", nil)
	}

	if response.StatusCode() != http.StatusOK {
		apiErr, err := resterrors.NewRestErrorFromBytes(response.Body())
		if err != nil {
			return nil, resterrors.NewInternalServerError("Invalid rest client error interface", "", err)
		}
		return nil, apiErr
	}

	var user User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, resterrors.NewInternalServerError("error unmarshalling", "", err)
	}
	return &user, nil
}
