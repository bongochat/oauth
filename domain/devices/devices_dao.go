package devices

import (
	"log"

	"github.com/bongochat/oauth/clients/cassandra"
	"github.com/bongochat/utils/resterrors"
)

const (
	queryGetDeviceList = "SELECT access_token, user_id, client_id, device_id, is_verified, device_type, device_model, ip, created_at FROM access_tokens WHERE user_id=? ALLOW FILTERING;"
)

func (r Devices) DeviceList(userId int64) ([]Devices, resterrors.RestError) {
	result := make([]Devices, 0)
	log.Println(userId)
	iter := cassandra.GetSession().Query(queryGetDeviceList, userId).Iter()
	defer iter.Close()
	for {
		var device Devices
		if !iter.Scan(&device.AccessToken, &device.UserId, &device.ClientId, &device.DeviceId, &device.IsVerified, &device.DeviceType, &device.DeviceModel, &device.IPAddress, &device.DateCreated) {
			break
		}
		log.Println(device)
		result = append(result, device)
	}
	if err := iter.Close(); err != nil {
		return nil, resterrors.NewInternalServerError("", "", err)
	}
	log.Println(result)
	return result, nil
}
