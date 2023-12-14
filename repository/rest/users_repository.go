package rest

import (
	"encoding/json"
	"time"

	"github.com/bongochat/bongochat-oauth/domain/users"
	"github.com/bongochat/bongochat-oauth/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRESTClient = rest.RequestBuilder{
		BaseURL: "https://api.users.bongo.chat",
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
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid rest client response")
	}

	if response.StatusCode > 299 {
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
