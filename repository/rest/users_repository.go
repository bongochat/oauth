package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bongochat/bongochat-oauth/domain/users"
	"github.com/bongochat/utils/logger"
	"github.com/bongochat/utils/resterrors"
	"github.com/go-resty/resty/v2"
)

var (
	userHostUrl     = os.Getenv("USER_HOST_URL")
	userLoginApiUrl = os.Getenv("USER_LOGIN_API_URL")
	userClient      = resty.New().SetBaseURL(userHostUrl)
)

type RESTUsersRepository interface {
	LoginUser(string, string) (*users.User, resterrors.RestError)
}

type usersRepository struct{}

func NewRepository() RESTUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(phone_number string, password string) (*users.User, resterrors.RestError) {
	request := users.UserLoginRequest{
		PhoneNumber: phone_number,
		Password:    password,
	}

	response, err := userClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(userLoginApiUrl)

	if err != nil {
		return nil, resterrors.NewInternalServerError("Could not login", err)
	}

	if response == nil {
		return nil, resterrors.NewInternalServerError("Invalid client response", nil)
	}
	logger.Info(fmt.Sprintf("RESP", response))

	if response.StatusCode() != http.StatusOK {
		apiErr, err := resterrors.NewRestErrorFromBytes(response.Body())
		logger.Error("Invalid error interface response", apiErr)
		logger.Error("Invalid error interface", err)
		if err != nil {
			return nil, resterrors.NewInternalServerError("Invalid rest client error interface", err)
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, resterrors.NewInternalServerError("error unmarshalling", err)
	}
	return &user, nil
}
