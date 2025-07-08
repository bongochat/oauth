package verify_token

import (
	"net/http"

	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/oauth/services"
	"github.com/bongochat/utils/resterrors"

	"github.com/gin-gonic/gin"
)

func VerifyAccessToken(c *gin.Context) {
	accessTokenString := c.Request.Header.Get("Authorization")
	if accessTokenString == "" {
		restErr := resterrors.NewBadRequestError("Invalid header information", "")
		c.JSON(http.StatusBadRequest, restErr)
		logger.RestErrorLog(restErr)
		return
	}
	accessTokenId := accessTokenString[len("Bearer "):]
	accessToken, err := services.TokenVerifyService.VerifyToken(accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		logger.RestErrorLog(err)
		return
	}
	c.JSON(http.StatusOK, accessToken.Marshall())
}

func VerifyClientAccessToken(c *gin.Context) {
	accessTokenString := c.Request.Header.Get("Authorization")
	if accessTokenString == "" {
		restErr := resterrors.NewBadRequestError("Invalid header information", "")
		c.JSON(http.StatusBadRequest, restErr)
		logger.RestErrorLog(restErr)
		return
	}
	accessTokenId := accessTokenString[len("Bearer "):]
	accessToken, err := services.TokenVerifyService.VerifyClientToken(accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		logger.RestErrorLog(err)
		return
	}
	c.JSON(http.StatusOK, accessToken.Marshall())
}
