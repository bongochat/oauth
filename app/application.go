package app

import (
	"net/http"

	"github.com/bongochat/bongochat-oauth/config"
	"github.com/bongochat/bongochat-oauth/handler"
	"github.com/bongochat/bongochat-oauth/repository/db"
	"github.com/bongochat/bongochat-oauth/repository/rest"
	"github.com/bongochat/bongochat-oauth/services/access_token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	conf   = config.GetConfig()
)

func StartApplication() {
	router.Use(cors.Default())
	atHandler := handler.NewHandler(access_token.NewService(rest.NewRepository(), db.NewRepository()))

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://bongo.chat")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connected",
			"status":  http.StatusOK,
		})
	})

	router.POST("/api/oauth/access-token/v1/", atHandler.CreateAccessToken)
	router.GET("/api/oauth/verify-token/v1/", atHandler.VerifyAccessToken)

	router.Run(conf.Port)
}
