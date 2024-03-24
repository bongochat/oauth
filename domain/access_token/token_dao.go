package access_token

import (
	"github.com/bongochat/oauth/clients/cassandra"
	"github.com/bongochat/utils/resterrors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, device_id, is_verified, device_type, device_model, ip, created_at FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, device_id, is_verified, device_type, device_model, ip, created_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);"
	queryDeleteAccessToken = "DELETE FROM access_tokens WHERE access_token=?"
)

func (r AccessToken) VerifyToken(userId int64, token string) (*AccessToken, resterrors.RestError) {
	var result AccessToken
	err := VerifyTokenString(token)
	if err != nil {
		return nil, resterrors.NewInternalServerError("Invalid access token", "", err)
	}
	if err := cassandra.GetSession().Query(queryGetAccessToken, token).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.DeviceId, &result.IsVerified, &result.DeviceType, &result.DeviceModel, &result.IPAddress, &result.DateCreated); err != nil {
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

func (at AccessToken) CreateToken() resterrors.RestError {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.DeviceId, at.IsVerified, at.DeviceType, at.DeviceModel, at.IPAddress, at.DateCreated).Exec(); err != nil {
		return resterrors.NewInternalServerError("Database error", "", err)
	}
	return nil
}

func (r AccessToken) DeleteToken(token string) resterrors.RestError {
	if err := cassandra.GetSession().Query(queryDeleteAccessToken, token).Exec(); err != nil {
		return resterrors.NewInternalServerError("Database error", "", err)
	}
	return nil
}
