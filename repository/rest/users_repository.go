package rest

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/bongochat/bongochat-oauth/domain/users"
	"github.com/bongochat/bongochat-oauth/logger"
	"github.com/bongochat/bongochat-oauth/utils/resterrors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	userHostUrl     = os.Getenv("USER_HOST_URL")
	userLoginApiUrl = os.Getenv("USER_LOGIN_API_URL")
	usersRESTClient = rest.RequestBuilder{
		BaseURL: userHostUrl,
		Timeout: 100 * time.Millisecond,
	}
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

	response := usersRESTClient.Post(userLoginApiUrl, request)
	if response == nil || response.Response == nil {
		logger.Error("Invalid response from user login", response.Err)
		return nil, resterrors.NewInternalServerError("Invalid client response", nil)
	}

	if response.StatusCode != http.StatusOK {
		apiErr, err := resterrors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, resterrors.NewInternalServerError("Invalid rest client error interface", err)
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, resterrors.NewInternalServerError("error unmarshalling", err)
	}
	return &user, nil
}
