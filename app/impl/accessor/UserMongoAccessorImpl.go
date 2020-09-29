package accessor

import "github.com/n4de4k/web-api-boilerplate/app/models"

type UserMongoAccessorImpl struct {
	// should has db attribute --> in this case mongo
}

func NewUserMongoAccessorImpl() *UserMongoAccessorImpl {
	return &UserMongoAccessorImpl{}
}

func(accessor *UserMongoAccessorImpl) GetUserByEmail(Email string) (*models.User, error) {
	return &models.User{
		Email:    Email,
		FullName: "Test User",
		Password: "Secret",
	},nil
}
