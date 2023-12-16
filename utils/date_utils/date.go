package date_utils

import "time"

const (
	dbDateLayout = "2006-01-02"
)

func GetNow() time.Time {
	return time.Now()
}

func GetCurrentDate() string {
	return GetNow().Format(dbDateLayout)
}
