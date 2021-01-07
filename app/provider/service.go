package provider

import (
	"github.com/n4de4k/web-api-boilerplate/app/service"
)

type ServiceComponent interface {
	GetUserService() service.UserService
}
