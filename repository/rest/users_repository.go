package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bongochat/bongochat-oauth/config"
	"github.com/bongochat/bongochat-oauth/domain/users"
	"github.com/bongochat/bongochat-oauth/utils/resterrors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	conf            = config.GetConfig()
	usersRESTClient = rest.RequestBuilder{
		BaseURL: conf.UserAPIBaseURL,
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

	response := usersRESTClient.Post(conf.UserLoginAPIURL, request)
	if response == nil || response.Response == nil {
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
