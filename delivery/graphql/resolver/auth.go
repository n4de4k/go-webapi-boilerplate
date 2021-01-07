package resolver

import "github.com/n4de4k/web-api-boilerplate/app/dto"

func (r *Resolver) SignIn(args *dto.SignInRequest) (*dto.SignInResponse, error) {
	user, err := r.Services.GetUserService().SignIn(args)
	if err != nil {
		return nil, err
	}

	return user, nil
}
