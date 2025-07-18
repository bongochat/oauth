package access_token

import (
	"context"
	"os"

	"github.com/bongochat/oauth/clients/mongodb"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/utils/resterrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r AccessToken) VerifyToken(token string) (*AccessToken, resterrors.RestError) {
	var result AccessToken
	err := VerifyTokenString(token)
	if err != nil {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Invalid access token", "", err)
	}
	filter := bson.M{"accesstoken": token}
	if err := mongodb.GetCollections().FindOne(context.Background(), filter).Decode(&result); err != nil {
		logger.ErrorLog(err)
		return nil, resterrors.NewUnauthorizedError("Access token not found", "")
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
		return nil, resterrors.NewUnauthorizedError("Access token not found", "")
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
	// testPhoneNumber := os.Getenv("TEST_PHONE_NUMBER")
	// // Determine the isVerified status based on the logic provided
	// if at.PhoneNumber == testPhoneNumber {
	// 	at.IsVerified = true
	// } else if existingToken.IsVerified {
	// 	at.IsVerified = true
	// } else {
	// 	at.IsVerified = false
	// }

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

func (at *AccessToken) GetToken() (*AccessToken, resterrors.RestError) {
	filter := bson.M{"accesstoken": at.AccessToken}

	// Fetch existing token to check current isVerified status
	existingToken := &AccessToken{}
	err := mongodb.GetCollections().FindOne(context.Background(), filter).Decode(existingToken)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Failed to retrieve existing document", "", err)
	}
	testPhoneNumber := os.Getenv("TEST_PHONE_NUMBER")
	// Determine the isVerified status based on the logic provided
	if at.PhoneNumber == testPhoneNumber {
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

func (r AccessToken) DeleteToken(token string) resterrors.RestError {
	filter := bson.M{"accesstoken": token}

	// Using DeleteOne to remove the document from the collection
	_, err := mongodb.GetCollections().DeleteOne(context.Background(), filter)
	if err != nil {
		logger.ErrorLog(err)
		return resterrors.NewInternalServerError("Database error", "", err)
	}
	return nil
}
