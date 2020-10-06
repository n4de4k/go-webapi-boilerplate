package impl

import (
	service2 "github.com/n4de4k/web-api-boilerplate/service"
	"github.com/n4de4k/web-api-boilerplate/service/impl"
)

type ServiceComponentImpl struct {
	userService service2.UserService
}

func NewServiceComponentImpl() *ServiceComponentImpl {
	dataComponent := NewDataComponentImpl()
	userService := impl.NewUserServiceImpl(dataComponent.GetUserAccessor())

	return &ServiceComponentImpl{
		userService,
	}
}

func (component *ServiceComponentImpl) GetUserService() service2.UserService {
	return component.userService
}
