package component

import (
	serviceImpl "github.com/n4de4k/web-api-boilerplate/app/impl/service"
	"github.com/n4de4k/web-api-boilerplate/app/service"
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
