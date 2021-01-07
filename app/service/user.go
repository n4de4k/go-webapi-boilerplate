package service

import (
	"github.com/n4de4k/web-api-boilerplate/app/dto"
)

type UserService interface {
	SignIn(request *dto.SignInRequest) (*dto.SignInResponse, error)
	GetUser(request *dto.GetUserRequest) (*dto.UserResponse, error)
}
