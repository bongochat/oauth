package http

import (
	"net/http"

	"github.com/bongochat/bongochat-oauth/domain/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByPhoneNumber(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetByPhoneNumber(c *gin.Context) {
	accessTokenId := c.Param("access_token")
	accessToken, err := handler.service.GetByPhoneNumber(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
