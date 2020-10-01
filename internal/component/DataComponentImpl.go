package component

import (
	accessorImpl "github.com/n4de4k/web-api-boilerplate/internal/accessor"
	"github.com/n4de4k/web-api-boilerplate/pkg/accessor"
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