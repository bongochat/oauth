package verify_device

import (
	"github.com/bongochat/oauth/clients/cassandra"
	"github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/utils/resterrors"
	"github.com/gocql/gocql"
)

const (
	queryVerifyDevice   = "UPDATE access_tokens SET is_verified=True WHERE access_token=?;"
	queryGetAccessToken = "SELECT access_token, user_id, client_id, device_id, is_verified, created_at FROM access_tokens WHERE access_token=?;"
)

func (r VerifyDevice) VerifyDevice(userId int64, token string) (*VerifyDevice, resterrors.RestError) {
	var result VerifyDevice
	err := access_token.VerifyTokenString(token)
	if err != nil {
		return nil, resterrors.NewInternalServerError("Invalid access token", "", err)
	}
	if err := cassandra.GetSession().Query(queryVerifyDevice, token).Exec(); err != nil {
		if err == gocql.ErrNotFound {
			return nil, resterrors.NewNotFoundError("Access token not found with the given phone number", "")
		}
		return nil, resterrors.NewInternalServerError("Database error", "", err)
	}
	if err := cassandra.GetSession().Query(queryGetAccessToken, token).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.DeviceId, &result.IsVerified, &result.DateCreated); err != nil {
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
