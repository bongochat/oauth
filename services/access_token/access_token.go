package access_token

import (
	"strings"
	"time"

	"github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/oauth/repository/rest"
	"github.com/bongochat/utils/resterrors"
)

type Service interface {
	VerifyToken(int64, string) (*access_token.AccessToken, resterrors.RestError)
	CreateToken(access_token.AccessTokenRequest) (*access_token.AccessToken, resterrors.RestError)
	DeleteToken(int64, string) resterrors.RestError
}

type service struct {
	restUsersRepo rest.RESTUsersRepository
	dbRepo        access_token.DBRepository
}

func NewService(usersRepo rest.RESTUsersRepository, dbRepo access_token.DBRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) VerifyToken(userId int64, accessTokenId string) (*access_token.AccessToken, resterrors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, resterrors.NewUnauthorizedError("Access token is required", "")
	}
	accessToken, err := s.dbRepo.VerifyToken(userId, accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) CreateToken(request access_token.AccessTokenRequest) (*access_token.AccessToken, resterrors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	if request.GrantType == "password" {
		user, err := s.restUsersRepo.LoginUser(request.PhoneNumber, request.Password)
		if err != nil {
			return nil, err
		}
		at := access_token.GetNewAccessToken(user.Id, request.DeviceId)
		// Generate a new access token:
		token, _ := at.Generate()
		at.AccessToken = token
		at.DateCreated = time.Now()

		// Save the new access token in Cassandra:
		if err := s.dbRepo.CreateToken(at); err != nil {
			return nil, err
		}
		return &at, nil
	} else {
		return nil, resterrors.NewBadRequestError("Invalid authentication type", request.GrantType)
	}

}

func (s *service) DeleteToken(userId int64, accessTokenId string) resterrors.RestError {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return resterrors.NewUnauthorizedError("Access token is required", "")
	}
	_, err := s.dbRepo.VerifyToken(userId, accessTokenId)
	if err != nil {
		return err
	}
	err = s.dbRepo.DeleteToken(accessTokenId)
	if err != nil {
		return err
	}
	return nil
}
