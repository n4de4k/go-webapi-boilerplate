package mongo

import (
	models2 "github.com/n4de4k/web-api-boilerplate/models"
)

type UserMongoAccessorImpl struct {
	// should has db attribute --> in this case mongo
}

func NewUserMongoAccessorImpl() *UserMongoAccessorImpl {
	return &UserMongoAccessorImpl{}
}

func(accessor *UserMongoAccessorImpl) GetUserByEmail(Email string) (*models2.User, error) {
	return &models2.User{
		Email:    Email,
		FullName: "Test User",
		Password: "Secret",
	},nil
}
