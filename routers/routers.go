package routers

import (
	"net/http"
	"os"
	"time"

	"github.com/bongochat/oauth/controllers/create_token"
	"github.com/bongochat/oauth/controllers/delete_token"
	"github.com/bongochat/oauth/controllers/verify_device"
	"github.com/bongochat/oauth/controllers/verify_token"
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

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://bongo.chat")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connected",
			"status":  http.StatusOK,
		})
	})

	tokenAPI := router.Group("/api/user")
	tokenAPI.POST("create-token/v1/", create_token.CreateAccessToken)
	tokenAPI.POST(":user_id/verify-device/v1/", verify_device.VerifyDevice)
	tokenAPI.GET(":user_id/verify-token/v1/", verify_token.VerifyAccessToken)
	tokenAPI.GET(":user_id/logout/v1/", delete_token.DeleteAccessToken)

	// run routes with port
	router.Run(os.Getenv("GO_PORT"))
}
