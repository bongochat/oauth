package services

import (
	"time"

	"github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/oauth/users"
	"github.com/bongochat/utils/resterrors"
)

var (
	TokenCreateService tokenCreateServiceInterface = &tokenCreateService{}
)

type tokenCreateService struct{}

type tokenCreateServiceInterface interface {
	CreateToken(access_token.AccessTokenRequest) (*access_token.AccessToken, resterrors.RestError)
}

func (s *tokenCreateService) CreateToken(request access_token.AccessTokenRequest) (*access_token.AccessToken, resterrors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	if request.GrantType == "password" {
		user, err := users.LoginUser(request.PhoneNumber, request.Password)
		if err != nil {
			logger.RestErrorLog(err)
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

		// Save the new access token in MongoDB:
		result, err := at.CreateToken()
		if err != nil {
			logger.RestErrorLog(err)
			return nil, err
		}
		return result, nil
	} else {
		return nil, resterrors.NewBadRequestError("Invalid authentication type", request.GrantType)
	}

}
