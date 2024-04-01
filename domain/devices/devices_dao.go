package devices

import (
	"context"

	"github.com/bongochat/oauth/clients/mongodb"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/utils/resterrors"
	"go.mongodb.org/mongo-driver/bson"
)

func (r Devices) DeviceList(userId int64) ([]Devices, resterrors.RestError) {
	result := make([]Devices, 0)

	filter := bson.M{"userid": userId}
	cursor, err := mongodb.GetCollections().Find(context.Background(), filter)
	if err != nil {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Error from device list query", "", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var device Devices
		if err := cursor.Decode(&device); err != nil {
			break
		} else {
			logger.ErrorLog(err)
		}
		result = append(result, device)
	}

	if err := cursor.Err(); err != nil {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Error from device list query", "", err)
	}
	return result, nil
}
