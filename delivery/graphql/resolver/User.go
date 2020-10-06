package resolver

import "github.com/n4de4k/web-api-boilerplate/models"

func (r *Resolver) GetUser(args *models.GetUserRequest) (*models.User, error) {
	user, err := r.Services.GetUserService().GetUser(args)
	if err != nil {
		return nil, err
	}

	return user, nil
}
