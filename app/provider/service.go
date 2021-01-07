package provider

import (
	"github.com/n4de4k/go-webapi-boilerplate/app/service"
)

type ServiceComponent interface {
	GetUserService() service.UserService
}
