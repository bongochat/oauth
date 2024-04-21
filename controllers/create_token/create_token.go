package create_token

import (
	"net/http"

	atDomain "github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/oauth/services"
	"github.com/bongochat/utils/resterrors"

	"github.com/gin-gonic/gin"
)

func CreateAccessToken(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := resterrors.NewBadRequestError("Invalid request", "")
		c.JSON(restErr.Status(), restErr)
		logger.RestErrorLog(restErr)
		return
	}

	accessToken, user, err := services.TokenCreateService.CreateToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		logger.RestErrorLog(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": accessToken.TokenMarshall(user),
		"status": http.StatusCreated,
	})
}
