package create_token

import (
	"fmt"
	"net/http"

	atDomain "github.com/bongochat/oauth/domain/access_token"
	"github.com/bongochat/oauth/logger"
	"github.com/bongochat/oauth/services"
	"github.com/bongochat/utils/resterrors"

	"github.com/gin-gonic/gin"
)

func CreateAccessToken(c *gin.Context) {
	var request atDomain.RegistrationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := resterrors.NewBadRequestError("Invalid request", "")
		fmt.Println(restErr)
		c.JSON(restErr.Status(), restErr)
		logger.RestErrorLog(restErr)
		return
	}
	fmt.Println(request.PhoneNumber)

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

func GetAccessToken(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := resterrors.NewBadRequestError("Invalid request", "")
		c.JSON(restErr.Status(), restErr)
		logger.RestErrorLog(restErr)
		return
	}
	fmt.Println(request.PhoneNumber)

	accessToken, user, err := services.TokenCreateService.GetToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		logger.RestErrorLog(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": accessToken.TokenMarshall(user),
		"status": http.StatusOK,
	})
}

func CreateClientAccessToken(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := resterrors.NewBadRequestError("Invalid request", "")
		c.JSON(restErr.Status(), restErr)
		logger.RestErrorLog(restErr)
		return
	}

	accessToken, client, err := services.TokenCreateService.CreateClientToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		logger.RestErrorLog(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": accessToken.ClientTokenMarshall(client),
		"status": http.StatusCreated,
	})
}
