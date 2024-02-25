package db

import (
	"github.com/bongochat/bongochat-oauth/clients/cassandra"
	"github.com/bongochat/bongochat-oauth/domain/access_token"
	"github.com/bongochat/utils/resterrors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, created_at FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, created_at) VALUES(?, ?, ?, ?);"
	queryDeleteAccessToken = "DELETE FROM access_tokens WHERE access_token=?"
)

func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	VerifyToken(int64, string) (*access_token.AccessToken, resterrors.RestError)
	CreateToken(access_token.AccessToken) resterrors.RestError
	DeleteToken(string) resterrors.RestError
}

type dbRepository struct {
}

func (r *dbRepository) VerifyToken(userId int64, token string) (*access_token.AccessToken, resterrors.RestError) {
	var result access_token.AccessToken
	err := access_token.VerifyToken(token)
	if err != nil {
		return nil, resterrors.NewInternalServerError("Invalid access token", "", err)
	}
	if err := cassandra.GetSession().Query(queryGetAccessToken, token).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.DateCreated); err != nil {
		if err == gocql.ErrNotFound {
			return nil, resterrors.NewNotFoundError("Access token not found with the given phone number", "")
		}
		return nil, resterrors.NewInternalServerError("Database error", "", err)
	}
	if userId != result.UserId {
		return nil, resterrors.NewUnauthorizedError("Access token not matching with the given user", "")
	}
	return &result, nil
}

func (r *dbRepository) CreateToken(at access_token.AccessToken) resterrors.RestError {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.DateCreated).Exec(); err != nil {
		return resterrors.NewInternalServerError("Database error", "", err)
	}
	return nil
}

func (r *dbRepository) DeleteToken(token string) resterrors.RestError {
	if err := cassandra.GetSession().Query(queryDeleteAccessToken, token).Exec(); err != nil {
		return resterrors.NewInternalServerError("Database error", "", err)
	}
	return nil
}
