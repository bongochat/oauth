package access_token

import (
	"context"
	"fmt"

	"github.com/bongochat/oauth/clients/mongodb"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/utils/resterrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (at AccessToken) CreateToken() (*AccessToken, resterrors.RestError) {
	_, err := mongodb.GetCollections().InsertOne(context.Background(), at)
	if err != nil {
		if wErr, ok := err.(mongo.WriteException); ok {
			for _, we := range wErr.WriteErrors {
				if we.Code == 11000 {
					at.IsVerified = true
					return &at, nil
				}
			}
		} else {
			logger.ErrorLog(err)
			return nil, resterrors.NewInternalServerError("Mongo Database error", "", err)
		}
		logger.ErrorLog(err)
		return nil, resterrors.NewInternalServerError("Mongo Database error", "", err)
	}
	return &at, nil
}

func (r AccessToken) DeleteToken(token string) resterrors.RestError {
	filter := bson.M{"accesstoken": token}
	result, err := mongodb.GetCollections().DeleteOne(context.Background(), filter)
	if err != nil {
		logger.ErrorLog(err)
		return resterrors.NewInternalServerError("Database error", "", err)
	}

	// Check if the deletion was successful
	if result.DeletedCount == 0 {
		logger.ErrorLog(err)
		return resterrors.NewInternalServerError("Token not found", "", err)
	}
	return nil
}
