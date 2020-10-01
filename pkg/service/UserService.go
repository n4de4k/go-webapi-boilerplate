package service

import v1 "github.com/n4de4k/web-api-boilerplate/api/resources/v1"

type UserService interface {
	SignIn(request v1.SignInRequest) (*v1.SignInResponse, error)
}
