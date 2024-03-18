package utils

import (
	"github.com/bongochat/utils/resterrors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, resterrors.RestError) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", resterrors.NewInternalServerError("Error form createing hash password", "", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(providedPassword string, hashedPassword string) resterrors.RestError {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	if err != nil {
		return resterrors.NewUnauthorizedError("Wrong password", "Wrong password")
	}
	return nil
}
