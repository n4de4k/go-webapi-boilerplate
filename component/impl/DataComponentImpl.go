package impl

import (
	accessor2 "github.com/n4de4k/web-api-boilerplate/accessor"
	"github.com/n4de4k/web-api-boilerplate/accessor/mongo"
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