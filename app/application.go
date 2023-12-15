package app

import (
	"github.com/bongochat/bongochat-oauth/http"
	"github.com/bongochat/bongochat-oauth/repository/db"
	"github.com/bongochat/bongochat-oauth/repository/rest"
	"github.com/bongochat/bongochat-oauth/services/access_token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(cors.Default())
	atHandler := http.NewHandler(access_token.NewService(rest.NewRepository(), db.NewRepository()))

	router.GET("/oauth/access-token/:access_token", atHandler.GetById)
	router.POST("/oauth/access-token/", atHandler.Create)

	router.Run(":8082")
}
