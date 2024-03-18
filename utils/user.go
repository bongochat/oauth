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
