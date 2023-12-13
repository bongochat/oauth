package db

import (
	"github.com/bongochat/bongochat-oauth/clients/cassandra"
	"github.com/bongochat/bongochat-oauth/domain/access_token"
	"github.com/bongochat/bongochat-oauth/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, phone_number, client_id, expires FROM access_token WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_token(access_token, phone_number, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateAccessToken = "UPDATE access_token SET expires=? WHERE access_token=?;"
)

func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetByPhoneNumber(string) (*access_token.AccessToken, *errors.RESTError)
	Create(*access_token.AccessToken) *errors.RESTError
	UpdateExpirationTime(*access_token.AccessToken) *errors.RESTError
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
	if err := session.Query(queryGetAccessToken, phoneNumber).Scan(&result.AccessToken, &result.PhoneNumber, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("Access token not found with the given phone number")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (r *dbRepository) Create(at *access_token.AccessToken) *errors.RESTError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken, at.AccessToken, at.PhoneNumber, at.ClientID, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at *access_token.AccessToken) *errors.RESTError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryUpdateAccessToken, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
