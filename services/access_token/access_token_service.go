package access_token

import (
	"strings"
	"time"

	"github.com/bongochat/bongochat-oauth/domain/access_token"
	"github.com/bongochat/bongochat-oauth/repository/db"
	"github.com/bongochat/bongochat-oauth/repository/rest"
	"github.com/bongochat/utils/resterrors"
)

type Service interface {
	VerifyToken(string) (*access_token.AccessToken, resterrors.RestError)
	CreateToken(access_token.AccessTokenRequest) (*access_token.AccessToken, resterrors.RestError)
}

type service struct {
	restUsersRepo rest.RESTUsersRepository
	dbRepo        db.DBRepository
}

func NewService(usersRepo rest.RESTUsersRepository, dbRepo db.DBRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) VerifyToken(accessTokenId string) (*access_token.AccessToken, resterrors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, resterrors.NewUnauthorizedError("Access token is required")
	}
	accessToken, err := s.dbRepo.VerifyToken(accessTokenId)
	if err != nil {
		return nil, resterrors.NewUnauthorizedError("Invalid access token")
	}
	return accessToken, nil
}

func (s *service) CreateToken(request access_token.AccessTokenRequest) (*access_token.AccessToken, resterrors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.restUsersRepo.LoginUser(request.PhoneNumber, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := access_token.GetNewAccessToken(user.Id)
	token, _ := at.Generate()
	at.AccessToken = token
	at.DateCreated = time.Now()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.CreateToken(at); err != nil {
		return nil, err
	}
	return &at, nil
}
