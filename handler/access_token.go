package handler

import (
	"net/http"

	atDomain "github.com/bongochat/bongochat-oauth/domain/access_token"
	"github.com/bongochat/bongochat-oauth/services/access_token"
	"github.com/bongochat/utils/resterrors"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	VerifyAccessToken(*gin.Context)
	CreateAccessToken(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) VerifyAccessToken(c *gin.Context) {
	accessTokenString := c.Request.Header.Get("Authorization")
	if accessTokenString == "" {
		restErr := resterrors.NewBadRequestError("Invalid header information")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	accessTokenId := accessTokenString[len("Bearer "):]
	accessToken, err := handler.service.VerifyToken(accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) CreateAccessToken(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := resterrors.NewBadRequestError("Invalid request")
		c.JSON(restErr.Status(), restErr)
		return
	}

	accessToken, err := handler.service.CreateToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
