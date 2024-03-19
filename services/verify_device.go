package services

import (
	"strings"

	"github.com/bongochat/oauth/domain/verify_device"
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
	accessTokenId := strings.TrimSpace(vd.AccessToken)
	if len(accessTokenId) == 0 {
		return nil, resterrors.NewUnauthorizedError("Access token is required", "")
	}
	accessToken, err := vd.VerifyDevice(userId, tokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
