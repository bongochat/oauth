package db

import (
	"github.com/bongochat/bongochat-oauth/clients/cassandra"
	"github.com/bongochat/bongochat-oauth/domain/access_token"
	"github.com/bongochat/bongochat-oauth/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, created_at FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, created_at) VALUES(?, ?, ?, ?);"
)

func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RESTError)
	Create(access_token.AccessToken) *errors.RESTError
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RESTError) {
	var result access_token.AccessToken
	err := access_token.VerifyToken(id)
	if err != nil {
		return nil, errors.NewInternalServerError("access token verification failed")
	}
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.DateCreated); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("Access token not found with the given phone number")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RESTError {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.DateCreated).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
