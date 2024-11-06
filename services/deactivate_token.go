package services

import (
	"strings"

	"github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/utils/resterrors"
)

var (
	TokenDeactivateService tokenDeactivateServiceInterface = &tokenDeactivateService{}
)

type tokenDeactivateService struct{}

type tokenDeactivateServiceInterface interface {
	DeactivateToken(int64, string) resterrors.RestError
}

func (s *tokenDeactivateService) DeactivateToken(userId int64, accessTokenId string) resterrors.RestError {
	at := &access_token.AccessToken{}
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		logger.ErrorMsgLog("Access token is required")
		return resterrors.NewUnauthorizedError("Access token is required", "")
	}
	_, err := TokenVerifyService.VerifyToken(userId, accessTokenId)
	if err != nil {
		logger.RestErrorLog(err)
		return err
	}
	err = at.DeactivateToken(accessTokenId)
	if err != nil {
		logger.RestErrorLog(err)
		return err
	}
	return nil
}
