package component

import (
	serviceImpl "github.com/n4de4k/web-api-boilerplate/internal/service"
	"github.com/n4de4k/web-api-boilerplate/pkg/service"
)

type ServiceComponentImpl struct {
	userService service.UserService
}

func NewServiceComponentImpl() *ServiceComponentImpl {
	dataComponent := NewDataComponentImpl()
	userService := serviceImpl.NewUserServiceImpl(dataComponent.GetUserAccessor())

	return &ServiceComponentImpl{
		userService,
	}
}

func (component *ServiceComponentImpl) GetUserService() service.UserService {
	return component.userService
}
