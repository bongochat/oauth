package services

import (
	"strings"

	"github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/utils/resterrors"
)

var (
	TokenDeleteService tokenDeleteServiceInterface = &tokenDeleteService{}
)

type tokenDeleteService struct{}

type tokenDeleteServiceInterface interface {
	DeleteToken(int64, string) resterrors.RestError
}

func (s *tokenDeleteService) DeleteToken(userId int64, accessTokenId string) resterrors.RestError {
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
	err = at.DeleteToken(accessTokenId)
	if err != nil {
		logger.RestErrorLog(err)
		return err
	}
	return nil
}
