package accessor

import (
	models2 "github.com/n4de4k/web-api-boilerplate/app/dto"
)

type UserAccessor interface {
	GetUserByEmail(Email string) (*models2.UserResponse, error)
}
