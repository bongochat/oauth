package services

import (
	"strings"
	"time"

	"github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/oauth/users"
	"github.com/bongochat/utils/resterrors"
)

var (
	TokenService tokenServiceInterface = &tokenService{}
)

type tokenService struct{}

type tokenServiceInterface interface {
	VerifyToken(int64, string) (*access_token.AccessToken, resterrors.RestError)
	CreateToken(access_token.AccessTokenRequest) (*access_token.AccessToken, resterrors.RestError)
	DeleteToken(int64, string) resterrors.RestError
}

func (service *tokenService) VerifyToken(userId int64, accessTokenId string) (*access_token.AccessToken, resterrors.RestError) {
	at := &access_token.AccessToken{}
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, resterrors.NewUnauthorizedError("Access token is required", "")
	}
	accessToken, err := at.VerifyToken(userId, accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *tokenService) CreateToken(request access_token.AccessTokenRequest) (*access_token.AccessToken, resterrors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	if request.GrantType == "password" {
		user, err := users.LoginUser(request.PhoneNumber, request.Password)
		if err != nil {
			return nil, err
		}
		at := access_token.GetNewAccessToken(user.Id, request.DeviceId)
		// Generate a new access token:
		token, _ := at.Generate()
		at.AccessToken = token
		at.DateCreated = time.Now()
		at.DeviceId = request.DeviceId
		at.DeviceType = request.DeviceType
		at.DeviceModel = request.DeviceModel
		at.IPAddress = request.IPAddress

		// Save the new access token in Cassandra:
		if err := at.CreateToken(); err != nil {
			return nil, err
		}
		return &at, nil
	} else {
		return nil, resterrors.NewBadRequestError("Invalid authentication type", request.GrantType)
	}

}

func (s *tokenService) DeleteToken(userId int64, accessTokenId string) resterrors.RestError {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return resterrors.NewUnauthorizedError("Access token is required", "")
	}
	_, err := s.VerifyToken(userId, accessTokenId)
	if err != nil {
		return err
	}
	err = s.DeleteToken(userId, accessTokenId)
	if err != nil {
		return err
	}
	return nil
}
