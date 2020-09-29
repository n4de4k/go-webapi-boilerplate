package accessor

import "github.com/n4de4k/web-api-boilerplate/app/models"

type UserAccessor interface {
	GetUserByEmail(Email string) (*models.User, error)
}
