package component

import (
	"github.com/n4de4k/web-api-boilerplate/service"
)

type ServiceComponent interface {
	GetUserService() service.UserService
}
