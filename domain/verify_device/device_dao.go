package verify_device

import (
	"context"

	"github.com/bongochat/oauth/clients/mongodb"
	"github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/utils/resterrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r VerifyDevice) VerifyDevice(userId int64, token string) (*VerifyDevice, resterrors.RestError) {
	var result VerifyDevice
	err := access_token.VerifyTokenString(token)
	if err != nil {
		return nil, resterrors.NewInternalServerError("Invalid access token", "", err)
	}

	filter := bson.M{"accesstoken": token}
	update := bson.M{
		"$set": bson.M{"isverified": true},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	if err := mongodb.GetCollections().FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&result); err != nil {
		return nil, resterrors.NewInternalServerError("Database error", "", err)
	}
	return &result, nil
}
