package utils

import (
	"strconv"

	"github.com/bongochat/utils/resterrors"
)

func GetUserID(userIdParam string) (int64, resterrors.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, resterrors.NewBadRequestError("user id should be a number", "")
	}
	return userId, nil
}

func GetClientID(clientIdParam string) (string, resterrors.RestError) {
	if clientIdParam == "" {
		return "", resterrors.NewBadRequestError("client id should be a string", "")
	}
	return clientIdParam, nil
}
