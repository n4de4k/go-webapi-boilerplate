package accessor

import (
	models2 "github.com/n4de4k/web-api-boilerplate/models"
)

type UserAccessor interface {
	GetUserByEmail(Email string) (*models2.User, error)
}
