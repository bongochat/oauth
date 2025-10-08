package verify_device

import (
	"context"
	"errors"

	"github.com/bongochat/oauth/clients/mongodb"
	"github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/utils/resterrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r VerifyDevice) VerifyDevice(accountNumber int64, token string) (*VerifyDevice, resterrors.RestError) {
	var result VerifyDevice
	err := access_token.VerifyTokenString(token)
	if err != nil {
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Invalid access token", "", err)
	}

	filter := bson.M{"accesstoken": token, "accountnumber": accountNumber}
	update := bson.M{
		"$set": bson.M{"isverified": true},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	if err := mongodb.GetCollections().FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, resterrors.NewUnauthorizedError("Access token is not valid", "Access token is not valid")
		}
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Database error", "", err)
	}
	return &result, nil
}
