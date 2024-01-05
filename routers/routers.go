package routers

import (
	"net/http"
	"os"

	"github.com/bongochat/bongochat-oauth/handler"
	"github.com/bongochat/bongochat-oauth/repository/db"
	"github.com/bongochat/bongochat-oauth/repository/rest"
	"github.com/bongochat/bongochat-oauth/services/access_token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func APIUrls() {
	corsconfig := cors.DefaultConfig()
	corsconfig.AllowOrigins = []string{"https://oauth.bongo.chat", "https://users.bongo.chat"}
	corsconfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	router.Use(cors.New(corsconfig))
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
	router.GET("/api/oauth/:user_id/verify-token/v1/", atHandler.VerifyAccessToken)

	// run routes with port
	router.Run(os.Getenv("GO_PORT"))
}
