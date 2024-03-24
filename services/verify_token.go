package services

import (
	"strings"

	"github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/utils/resterrors"
)

var (
	TokenVerifyService tokenVerifyServiceInterface = &tokenVerifyService{}
)

type tokenVerifyService struct{}

type tokenVerifyServiceInterface interface {
	VerifyToken(int64, string) (*access_token.AccessToken, resterrors.RestError)
}

func (service *tokenVerifyService) VerifyToken(userId int64, accessTokenId string) (*access_token.AccessToken, resterrors.RestError) {
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
