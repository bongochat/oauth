package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now()
}

func GetCurrentDate() string {
	return GetNow().Format(apiDateLayout)
}
