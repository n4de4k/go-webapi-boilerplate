package service

import (
	"github.com/n4de4k/web-api-boilerplate/models"
)

type UserService interface {
	SignIn(request *models.SignInRequest) (*models.SignInResponse, error)
	GetUser(request *models.GetUserRequest) (*models.User, error)
}
