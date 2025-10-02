package delete_token

import (
	"net/http"

	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/oauth/services"
	"github.com/bongochat/utils/resterrors"

	"github.com/gin-gonic/gin"
)

func DeleteAccessToken(c *gin.Context) {
	accessTokenString := c.Request.Header.Get("Authorization")
	if accessTokenString == "" {
		restErr := resterrors.NewBadRequestError("Invalid header information", "")
		c.JSON(http.StatusBadRequest, restErr)
		logger.RestErrorLog(restErr)
		return
	}
	accessTokenId := accessTokenString[len("Bearer "):]
	_, err := services.TokenVerifyService.VerifyToken(accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		logger.RestErrorLog(err)
		return
	}
	err = services.TokenDeleteService.DeleteToken(accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		logger.RestErrorLog(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Device Logout successfully.",
		"status":  http.StatusOK,
	})
}
