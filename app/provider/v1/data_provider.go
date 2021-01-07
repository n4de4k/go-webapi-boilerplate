package v1

import (
	accessor2 "github.com/n4de4k/web-api-boilerplate/app/accessor"
	"github.com/n4de4k/web-api-boilerplate/app/accessor/mongo"
)

type DataComponentImpl struct {
	userAccessor accessor2.UserAccessor
}

func NewDataComponentImpl() *DataComponentImpl {
	userAccessor := mongo.NewUserMongoAccessorImpl()


	return &DataComponentImpl{
		userAccessor,
	}
}

func (component DataComponentImpl) GetUserAccessor() accessor2.UserAccessor {
	return component.userAccessor
}