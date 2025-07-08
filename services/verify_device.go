package services

import (
	"strings"

	"github.com/bongochat/oauth/domain/verify_device"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/utils/resterrors"
)

var (
	DeviceService deviceServiceInterface = &deviceService{}
)

type deviceService struct{}

type deviceServiceInterface interface {
	VerifyDevice(int64, string) (*verify_device.VerifyDevice, resterrors.RestError)
}

func (service *deviceService) VerifyDevice(userId int64, tokenId string) (*verify_device.VerifyDevice, resterrors.RestError) {
	vd := &verify_device.VerifyDevice{}
	accessTokenId := strings.TrimSpace(tokenId)
	if len(accessTokenId) == 0 {
		logger.ErrorMsgLog("Access token is required")
		return nil, resterrors.NewUnauthorizedError("Access token is required", "")
	}
	_, err := TokenVerifyService.VerifyToken(accessTokenId)
	if err != nil {
		logger.RestErrorLog(err)
		return nil, err
	}
	accessToken, err := vd.VerifyDevice(userId, tokenId)
	if err != nil {
		logger.RestErrorLog(err)
		return nil, err
	}
	return accessToken, nil
}
