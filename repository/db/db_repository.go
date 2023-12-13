package db

import (
	"github.com/bongochat/bongochat-oauth/clients/cassandra"
	"github.com/bongochat/bongochat-oauth/domain/access_token"
	"github.com/bongochat/bongochat-oauth/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken = "SELECT access_token, phone_number, client_id, expires FROM access_token WHERE access_token=?;"
)

func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetByPhoneNumber(string) (*access_token.AccessToken, *errors.RESTError)
}

type dbRepository struct {
}

func (r *dbRepository) GetByPhoneNumber(phoneNumber string) (*access_token.AccessToken, *errors.RESTError) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}

	defer session.Close()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, phoneNumber).Scan(&result.AccessToken, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("Access token not found with the given phone number")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}
