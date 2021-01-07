package impl

import (
	"errors"
	accessor2 "github.com/n4de4k/go-webapi-boilerplate/app/accessor"
	"github.com/n4de4k/go-webapi-boilerplate/app/dto"
)

type UserServiceImpl struct {
	userAccessor accessor2.UserAccessor
}

func NewUserServiceImpl(_userAccessor accessor2.UserAccessor) *UserServiceImpl {
	return &UserServiceImpl{_userAccessor}
}

func (impl *UserServiceImpl) SignIn(request *dto.SignInRequest) (*dto.SignInResponse, error) {
	user, err := impl.userAccessor.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if user.Password != request.Password {
		return nil, errors.New("Invalid Password")
	}

	return &dto.SignInResponse{Token: user.FullName + user.Email}, nil
}

func (impl *UserServiceImpl) GetUser(request *dto.GetUserRequest) (*dto.UserResponse, error) {
	user, err := impl.userAccessor.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
