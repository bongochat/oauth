package routers

import (
	"net/http"
	"os"
	"time"

	"github.com/bongochat/oauth/controllers/create_token"
	"github.com/bongochat/oauth/controllers/delete_token"
	"github.com/bongochat/oauth/controllers/devices"
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
	corsconfig.AllowOrigins = []string{"https://*.bongo.chat"}
	corsconfig.AllowMethods = []string{"GET", "POST"}
	corsconfig.AllowHeaders = []string{"Origin", "Content-Type", "Cache-Control", "max-age=3600, private"}
	corsconfig.AllowCredentials = true
	corsconfig.ExposeHeaders = []string{"Authorization", "Content-Length"}
	corsconfig.MaxAge = 12 * time.Hour

	router.Use(cors.New(corsconfig))

	router.Use(func(c *gin.Context) {
		// Set cache control headers
		c.Header("Cache-Control", "max-age=3600, public") // Cache response for 1 hour
		c.Next()
	})

	router.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://bongo.chat")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connected",
			"status":  http.StatusOK,
		})
	})

	tokenAPI := router.Group("/api/v1/user")
	tokenAPI.POST("create-token/", create_token.CreateAccessToken)
	tokenAPI.POST(":user_id/verify-device/", verify_device.VerifyDevice)
	tokenAPI.GET(":user_id/verify-token/", verify_token.VerifyAccessToken)
	tokenAPI.GET(":user_id/logout/", delete_token.DeleteAccessToken)
	tokenAPI.GET(":user_id/device-list/", devices.DeviceList)

	clientTokenAPI := router.Group("/api/v1/client")
	clientTokenAPI.POST("create-token/", create_token.CreateClientAccessToken)
	clientTokenAPI.GET(":client_id/verify-token/", verify_token.VerifyClientAccessToken)

	// run routes with port
	router.Run(os.Getenv("GO_PORT"))
}
