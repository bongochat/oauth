package routers

import (
	"net/http"
	"os"
	"time"

	token "github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/oauth/handler"
	"github.com/bongochat/oauth/repository/rest"
	"github.com/bongochat/oauth/services/access_token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func APIUrls() {
	corsconfig := cors.DefaultConfig()
	corsconfig.AllowOrigins = []string{"*"}
	corsconfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsconfig.AllowHeaders = []string{"Origin"}
	corsconfig.AllowCredentials = true
	corsconfig.ExposeHeaders = []string{"Authorization", "Content-Length"}
	corsconfig.MaxAge = 12 * time.Hour

	router.Use(cors.New(corsconfig))
	atHandler := handler.NewHandler(access_token.NewService(rest.NewRepository(), token.NewRepository()))

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
	router.GET("/api/oauth/:user_id/logout/v1/", atHandler.DeleteAccessToken)

	// run routes with port
	router.Run(os.Getenv("GO_PORT"))
}
