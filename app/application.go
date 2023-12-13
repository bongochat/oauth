package app

import (
	"github.com/bongochat/bongochat-oauth/domain/access_token"
	"github.com/bongochat/bongochat-oauth/http"
	"github.com/bongochat/bongochat-oauth/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access-token/:access_token", atHandler.GetByPhoneNumber)
	router.POST("/oauth/access-token/", atHandler.Create)

	router.Run(":8082")
}
