package services

import (
	"strings"

	"github.com/bongochat/oauth/domain/devices"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/utils/resterrors"
)

var (
	DeviceListService deviceListServiceInterface = &deviceListService{}
)

type deviceListService struct{}

type deviceListServiceInterface interface {
	DeviceList(int64, string) ([]devices.Devices, resterrors.RestError)
}

func (service *deviceListService) DeviceList(userId int64, tokenId string) ([]devices.Devices, resterrors.RestError) {
	devices := &devices.Devices{}
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
	deviceList, err := devices.DeviceList(userId)
	if err != nil {
		logger.RestErrorLog(err)
		return nil, err
	}
	return deviceList, nil
}
