package logger

import (
	"fmt"
	"log"

	"github.com/bongochat/utils/resterrors"
	"github.com/getsentry/sentry-go"
)

func InitSentry(dsn string, env string, release string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         dsn,
		Environment: env,
		Release:     release,
	})
	if err != nil {
		log.Fatalf("sentry.Init failed: %v", err)
	}
}

// InfoLog writes an info message
func InfoLog(message string) {
	sentry.CaptureMessage("[INFO] " + message)
}

// ErrorMsgLog logs error messages
func ErrorMsgLog(message string) {
	sentry.CaptureMessage("[INFO] " + message)
}

// ErrorLog logs error objects
func ErrorLog(err error) {
	sentry.CaptureException(err)
}

// RestErrorLog logs custom RestError
func RestErrorLog(err resterrors.RestError) {
	fmt.Println("ERROR LOG: ", err)
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetTag("type", "rest_error")
		scope.SetExtra("status", err.Status())
		scope.SetExtra("message", err.Message())
		scope.SetExtra("error", err.Error())
		sentry.CaptureException(err)
	})
}
