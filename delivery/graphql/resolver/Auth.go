package resolver

import "github.com/n4de4k/web-api-boilerplate/models"

func (r *Resolver) SignIn(args *models.SignInRequest) (*models.SignInResponse, error) {
	user, err := r.Services.GetUserService().SignIn(args)
	if err != nil {
		return nil, err
	}

	return user, nil
}
