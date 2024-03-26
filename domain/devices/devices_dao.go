package devices

import (
	"context"
	"log"

	"github.com/bongochat/oauth/clients/mongodb"
	"github.com/bongochat/utils/resterrors"
	"go.mongodb.org/mongo-driver/bson"
)

func (r Devices) DeviceList(userId int64) ([]Devices, resterrors.RestError) {
	result := make([]Devices, 0)
	log.Println(userId)

	filter := bson.M{"userid": userId}
	cursor, err := mongodb.GetCollections().Find(context.Background(), filter)
	if err != nil {
		return nil, resterrors.NewInternalServerError("Error from device list query", "", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var device Devices
		if err := cursor.Decode(&device); err != nil {
			break
		}
		result = append(result, device)
	}

	if err := cursor.Err(); err != nil {
		return nil, resterrors.NewInternalServerError("Error from device list query", "", err)
	}
	return result, nil
}
