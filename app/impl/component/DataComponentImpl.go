package component

import (
	"github.com/n4de4k/web-api-boilerplate/app/accessor"
	accessorImpl "github.com/n4de4k/web-api-boilerplate/app/impl/accessor"
)

type DataComponentImpl struct {
	userAccessor accessor.UserAccessor
}

func NewDataComponentImpl() *DataComponentImpl {
	userAccessor := accessorImpl.NewUserMongoAccessorImpl()

	return &DataComponentImpl{
		userAccessor,
	}
}

func (component DataComponentImpl) GetUserAccessor() accessor.UserAccessor {
	return component.userAccessor
}