package utils

import (
	"strconv"

	"github.com/bongochat/utils/resterrors"
)

func ValidateAccountNumber(accountNumberParam string) (int64, resterrors.RestError) {
	accountNumber, userErr := strconv.ParseInt(accountNumberParam, 10, 64)
	if userErr != nil {
		return 0, resterrors.NewBadRequestError("user id should be a number", "")
	}
	return accountNumber, nil
}

func ValidateClientID(clientIdParam string) (string, resterrors.RestError) {
	if clientIdParam == "" {
		return "", resterrors.NewBadRequestError("client id should be a string", "")
	}
	return clientIdParam, nil
}
