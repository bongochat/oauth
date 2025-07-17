package main

import (
	"log"
	"os"
	"time"

	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/oauth/routers"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		logger.ErrorMsgLog("Error loading .env file")
	}

	// configure sentry log
	environment := os.Getenv("ENV")
	sentry_url := os.Getenv("SENTRY_URL")
	logger.InitSentry(sentry_url, environment, "v1.0.0")
	defer sentry.Flush(2 * time.Second)

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	logger.InfoLog("OAuth server started...")
	// load endpoints
	routers.APIUrls()
}
