package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/n4de4k/go-webapi-boilerplate/app/dto"
	"net/http"
)

func (handler *Handler) SignIn(c *gin.Context) {
	var request dto.SignInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	response, err := handler.serviceComponent.GetUserService().SignIn(&request)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}