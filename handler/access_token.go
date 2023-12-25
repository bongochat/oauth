package handler

import (
	"net/http"
	"strconv"

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

func getUserId(userIdParam string) (int64, resterrors.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, resterrors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func (handler *accessTokenHandler) VerifyAccessToken(c *gin.Context) {
	accessTokenString := c.Request.Header.Get("Authorization")
	if accessTokenString == "" {
		restErr := resterrors.NewBadRequestError("Invalid header information")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	userId, userIdErr := getUserId(c.Param("user_id"))
	if userIdErr != nil {
		c.JSON(userIdErr.Status(), userIdErr)
		return
	}
	accessTokenId := accessTokenString[len("Bearer "):]
	accessToken, err := handler.service.VerifyToken(userId, accessTokenId)
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
