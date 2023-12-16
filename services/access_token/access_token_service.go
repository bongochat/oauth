package access_token

import (
	"strings"

	"github.com/bongochat/bongochat-oauth/domain/access_token"
	"github.com/bongochat/bongochat-oauth/repository/db"
	"github.com/bongochat/bongochat-oauth/repository/rest"
	"github.com/bongochat/bongochat-oauth/utils/date_utils"
	"github.com/bongochat/bongochat-oauth/utils/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RESTError)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RESTError)
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

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RESTError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RESTError) {
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
	token, err := at.Generate()
	if err != nil {
		errors.NewInternalServerError("Generate access token failed")
	}
	at.AccessToken = token
	at.DateCreated = date_utils.GetCurrentDate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}
