package mongo

import (
	models2 "github.com/n4de4k/web-api-boilerplate/app/dto"
)

type UserMongoAccessorImpl struct {
	// should has db attribute --> in this case mongo
}

func NewUserMongoAccessorImpl() *UserMongoAccessorImpl {
	return &UserMongoAccessorImpl{}
}

func(accessor *UserMongoAccessorImpl) GetUserByEmail(Email string) (*models2.UserResponse, error) {
	return &models2.UserResponse{
		Email:    Email,
		FullName: "Test UserResponse",
		Password: "Secret",
	},nil
}
