package service

import (
	"errors"
	v1 "github.com/n4de4k/web-api-boilerplate/api/resources/v1"
	"github.com/n4de4k/web-api-boilerplate/pkg/accessor"
)

type UserServiceImpl struct {
	userAccessor accessor.UserAccessor
}

func NewUserServiceImpl(_userAccessor accessor.UserAccessor) *UserServiceImpl {
	return &UserServiceImpl{_userAccessor}
}

func (impl *UserServiceImpl) SignIn(request v1.SignInRequest) (*v1.SignInResponse, error) {
	user, err := impl.userAccessor.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if user.Password != request.Password {
		return nil, errors.New("Invalid Password")
	}

	return &v1.SignInResponse{Token: user.FullName + user.Email}, nil
}
