package logger

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/logging"
	"github.com/bongochat/utils/resterrors"
)

var Client *logging.Client

var projectID = os.Getenv("PROJECT_ID")
var logName = os.Getenv("LOG_NAME")

func init() {
	ctx := context.Background()

	var err error
	Client, err = logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

}

func InfoLog(message string) {
	logger := Client.Logger(logName).StandardLogger(logging.Info)
	logger.Println(message)
}

func ErrorMsgLog(message string) {
	logger := Client.Logger(logName).StandardLogger(logging.Error)
	logger.Println(message)
}

func ErrorLog(err error) {
	logger := Client.Logger(logName).StandardLogger(logging.Error)
	logger.Println(err)
}

func RestErrorLog(err resterrors.RestError) {
	logger := Client.Logger(logName).StandardLogger(logging.Error)
	logger.Println(err)
}

func Close() {
	if Client != nil {
		if err := Client.Close(); err != nil {
			log.Fatalf("Failed to close client: %v", err)
		}
	}
}
