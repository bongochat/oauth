package main

import (
	"log"
	"os"

	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/oauth/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		logger.ErrorMsgLog("Error loading .env file")
	}

	// logger.InfoLog("Starting oauth server")

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	logger.InfoLog("OAuth server started...")
	// load endpoints
	routers.APIUrls()

	// close the logging channel
	logger.Close()
}
