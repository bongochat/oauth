package services

import (
	"strings"

	"github.com/bongochat/oauth/domain/devices"
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
		return nil, resterrors.NewUnauthorizedError("Access token is required", "")
	}
	_, err := TokenVerifyService.VerifyToken(userId, accessTokenId)
	if err != nil {
		return nil, err
	}
	deviceList, err := devices.DeviceList(userId)
	if err != nil {
		return nil, err
	}
	return deviceList, nil
}
