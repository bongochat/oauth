package users

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/bongochat/utils/resterrors"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	ClientName   string `json:"client_name"`
	DateCreated  string `json:"date_created"`
	IsActive     bool   `json:"is_active"`
}

type LoginRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

var (
	clientHostUrl     = os.Getenv("USER_HOST_URL")
	clientLoginApiUrl = os.Getenv("USER_LOGIN_API_URL")
	clientClient      = resty.New().SetBaseURL(clientHostUrl)
)

func LoginClient(clientId string, clientSecret string) (*Client, resterrors.RestError) {
	request := LoginRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}

	response, err := clientClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(clientLoginApiUrl)

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

	var client Client
	if err := json.Unmarshal(response.Body(), &client); err != nil {
		return nil, resterrors.NewInternalServerError("error unmarshalling", "", err)
	}
	return &client, nil
}
