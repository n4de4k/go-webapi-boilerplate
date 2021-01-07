package resolver

import "github.com/n4de4k/web-api-boilerplate/app/dto"

func (r *Resolver) GetUser(args *dto.GetUserRequest) (*dto.UserResponse, error) {
	user, err := r.Services.GetUserService().GetUser(args)
	if err != nil {
		return nil, err
	}

	return user, nil
}
