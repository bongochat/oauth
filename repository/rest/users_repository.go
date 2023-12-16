package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bongochat/bongochat-oauth/config"
	"github.com/bongochat/bongochat-oauth/domain/users"
	"github.com/bongochat/bongochat-oauth/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	conf            = config.GetConfig()
	usersRESTClient = rest.RequestBuilder{
		BaseURL: conf.UserAPIURL,
		Timeout: 100 * time.Millisecond,
	}
)

type RESTUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RESTError)
}

type usersRepository struct{}

func NewRepository() RESTUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(phone_number string, password string) (*users.User, *errors.RESTError) {
	request := users.UserLoginRequest{
		PhoneNumber: phone_number,
		Password:    password,
	}

	response := usersRESTClient.Post("/users/login/", request)
	log.Println(response, conf.UserAPIURL)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid rest client response")
	}

	if response.StatusCode != http.StatusOK {
		var restErr errors.RESTError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("Invalid rest client error interface")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error unmarshalling")
	}
	return &user, nil
}
