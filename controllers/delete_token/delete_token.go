package delete_token

import (
	"net/http"

	"github.com/bongochat/oauth/services"
	"github.com/bongochat/oauth/utils"
	"github.com/bongochat/utils/resterrors"

	"github.com/gin-gonic/gin"
)

func DeleteAccessToken(c *gin.Context) {
	accessTokenString := c.Request.Header.Get("Authorization")
	if accessTokenString == "" {
		restErr := resterrors.NewBadRequestError("Invalid header information", "")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	userId, userIdErr := utils.GetUserID(c.Param("user_id"))
	if userIdErr != nil {
		c.JSON(userIdErr.Status(), userIdErr)
		return
	}
	accessTokenId := accessTokenString[len("Bearer "):]
	_, err := services.TokenService.VerifyToken(userId, accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	err = services.TokenService.DeleteToken(userId, accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successfully.",
		"status":  http.StatusOK,
	})
}
