package main

import (
	"github.com/bongochat/bongochat-oauth/app"
	"github.com/bongochat/bongochat-oauth/logger"
)

func main() {
	app.StartApplication()
	logger.Info("Starting application with credentials ")
}
