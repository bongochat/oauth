package access_token

import (
	"log"
	"strings"

	"github.com/bongochat/bongochat-oauth/utils/errors"
)

type Repository interface {
	GetByPhoneNumber(string) (*AccessToken, *errors.RESTError)
	Create(*AccessToken) *errors.RESTError
	UpdateExpirationTime(*AccessToken) *errors.RESTError
}

type Service interface {
	GetByPhoneNumber(string) (*AccessToken, *errors.RESTError)
	Create(*AccessToken) *errors.RESTError
	UpdateExpirationTime(*AccessToken) *errors.RESTError
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByPhoneNumber(accessTokenId string) (*AccessToken, *errors.RESTError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	log.Println(accessTokenId, "access token")
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token Id")
	}
	log.Println(accessTokenId)
	accessToken, err := s.repository.GetByPhoneNumber(accessTokenId)
	if err != nil {
		return nil, err
	}
	log.Println(accessToken, "AT")
	return accessToken, nil
}

func (s *service) Create(at *AccessToken) *errors.RESTError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at *AccessToken) *errors.RESTError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
