package db

import (
	"log"

	"github.com/bongochat/bongochat-oauth/clients/cassandra"
	"github.com/bongochat/bongochat-oauth/domain/access_token"
	"github.com/bongochat/utils/resterrors"
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
	VerifyToken(string) (*access_token.AccessToken, resterrors.RestError)
	CreateToken(access_token.AccessToken) resterrors.RestError
}

type dbRepository struct {
}

func (r *dbRepository) VerifyToken(token string) (*access_token.AccessToken, resterrors.RestError) {
	var result access_token.AccessToken
	log.Println(token, "GET TOKEN")
	err := access_token.VerifyToken(token)
	log.Println(err)
	if err != nil {
		return nil, resterrors.NewInternalServerError("Invalid access token", err)
	}
	if err := cassandra.GetSession().Query(queryGetAccessToken, token).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.DateCreated); err != nil {
		if err == gocql.ErrNotFound {
			return nil, resterrors.NewNotFoundError("Access token not found with the given phone number")
		}
		return nil, resterrors.NewInternalServerError("Database error", err)
	}
	log.Println(result)
	return &result, nil
}

func (r *dbRepository) CreateToken(at access_token.AccessToken) resterrors.RestError {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.DateCreated).Exec(); err != nil {
		return resterrors.NewInternalServerError("Database error", err)
	}
	return nil
}
