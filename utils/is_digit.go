package utils

import (
	"strconv"
)

func IsDigit(v string) bool {
	if _, err := strconv.ParseInt(v, 10, 64); err == nil {
		return true
	}
	return false
}
