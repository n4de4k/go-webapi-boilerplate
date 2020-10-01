package rest

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/n4de4k/web-api-boilerplate/api/resources/v1"
	"net/http"
)

func (handler *Handler) SignIn(c *gin.Context) {
	var request v1.SignInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	response, err := handler.serviceComponent.GetUserService().SignIn(request)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}