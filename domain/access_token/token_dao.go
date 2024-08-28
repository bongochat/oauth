package access_token

import (
	"context"
	"fmt"

	"github.com/bongochat/oauth/clients/mongodb"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/utils/resterrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r AccessToken) VerifyToken(userId int64, token string) (*AccessToken, resterrors.RestError) {
	var result AccessToken
	err := VerifyTokenString(token)
	if err != nil {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Invalid access token", "", err)
	}
	filter := bson.M{"accesstoken": token}
	if err := mongodb.GetCollections().FindOne(context.Background(), filter).Decode(&result); err != nil {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Access token not found", "", err)
	}
	if userId != result.UserId {
		logger.ErrorMsgLog(fmt.Sprintf("Access token not matching with the given user=%d", result.UserId))
		return nil, resterrors.NewUnauthorizedError("Access token not matching with the given user", "")
	}
	return &result, nil
}

func (r AccessToken) VerifyClientToken(token string) (*AccessToken, resterrors.RestError) {
	var result AccessToken
	err := VerifyTokenString(token)
	if err != nil {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Invalid access token", "", err)
	}
	filter := bson.M{"accesstoken": token}
	if err := mongodb.GetCollections().FindOne(context.Background(), filter).Decode(&result); err != nil {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Access token not found", "", err)
	}
	return &result, nil
}

func (at *AccessToken) CreateToken() (*AccessToken, resterrors.RestError) {
	filter := bson.M{"accesstoken": at.AccessToken}

	// Fetch existing token to check current isVerified status
	existingToken := &AccessToken{}
	err := mongodb.GetCollections().FindOne(context.Background(), filter).Decode(existingToken)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Failed to retrieve existing document", "", err)
	}

	// Determine the isVerified status based on the logic provided
	if at.PhoneNumber == "8801200000000" {
		at.IsVerified = true
	} else if existingToken.IsVerified {
		at.IsVerified = true
	} else {
		at.IsVerified = false
	}

	update := bson.M{
		"$set": bson.M{
			"userid":       at.UserId,
			"clientid":     at.ClientId,
			"clientsecret": at.ClientSecret,
			"deviceid":     at.DeviceId,
			"devicetype":   at.DeviceType,
			"devicemodel":  at.DeviceModel,
			"ipaddress":    at.IPAddress,
			"isactive":     at.IsActive,
			"isverified":   at.IsVerified,
			"datecreated":  at.CreatedAt,
			"dateupdated":  at.UpdatedAt,
		},
	}

	options := options.Update().SetUpsert(true)
	result, err := mongodb.GetCollections().UpdateOne(context.Background(), filter, update, options)
	if err != nil {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Mongo Database error", "", err)
	}

	if result.ModifiedCount > 0 {
		err := mongodb.GetCollections().FindOne(context.Background(), filter).Decode(at)
		if err != nil {
			logger.ErrorLog(err)
			return nil, resterrors.NewInternalServerError("Failed to retrieve updated document", "", err)
		}
	}

	return at, nil
}

func (at *AccessToken) CreateClientToken() (*AccessToken, resterrors.RestError) {
	filter := bson.M{"accesstoken": at.AccessToken}
	update := bson.M{
		"$set": bson.M{
			"userid":       at.UserId,
			"clientid":     at.ClientId,
			"clientsecret": at.ClientSecret,
			"deviceid":     at.DeviceId,
			"devicetype":   at.DeviceType,
			"devicemodel":  at.DeviceModel,
			"ipaddress":    at.IPAddress,
			"isactive":     at.IsActive,
			"isverified":   at.IsVerified,
			"datecreated":  at.CreatedAt,
			"dateupdated":  at.UpdatedAt,
		},
	}
	options := options.Update().SetUpsert(true)
	result, err := mongodb.GetCollections().UpdateOne(context.Background(), filter, update, options)
	if err != nil {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Mongo Database error", "", err)
	}

	if result.ModifiedCount > 0 {
		err := mongodb.GetCollections().FindOne(context.Background(), filter).Decode(at)
		if err != nil {
			logger.ErrorLog(err)
			return nil, resterrors.NewInternalServerError("Failed to retrieve updated document", "", err)
		}
	}

	return at, nil
}

func (r AccessToken) DeactivateToken(token string) resterrors.RestError {
	filter := bson.M{"accesstoken": token}
	update := bson.M{"$set": bson.M{"isactive": false}}
	_, err := mongodb.GetCollections().UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.ErrorLog(err)
		return resterrors.NewInternalServerError("Database err", "", err)
	}
	return nil
}
