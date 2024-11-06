package deactivate_token

import (
	"net/http"

	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/oauth/services"
	"github.com/bongochat/oauth/utils"
	"github.com/bongochat/utils/resterrors"

	"github.com/gin-gonic/gin"
)

func DeactivateAccessToken(c *gin.Context) {
	accessTokenString := c.Request.Header.Get("Authorization")
	if accessTokenString == "" {
		restErr := resterrors.NewBadRequestError("Invalid header information", "")
		c.JSON(http.StatusBadRequest, restErr)
		logger.RestErrorLog(restErr)
		return
	}
	userId, userIdErr := utils.GetUserID(c.Param("user_id"))
	if userIdErr != nil {
		c.JSON(userIdErr.Status(), userIdErr)
		logger.RestErrorLog(userIdErr)
		return
	}
	accessTokenId := accessTokenString[len("Bearer "):]
	_, err := services.TokenVerifyService.VerifyToken(userId, accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		logger.RestErrorLog(err)
		return
	}
	err = services.TokenDeactivateService.DeactivateToken(userId, accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		logger.RestErrorLog(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successfully.",
		"status":  http.StatusOK,
	})
}
