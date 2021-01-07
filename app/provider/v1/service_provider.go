package v1

import (
	service2 "github.com/n4de4k/go-webapi-boilerplate/app/service"
	"github.com/n4de4k/go-webapi-boilerplate/app/service/impl"
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
