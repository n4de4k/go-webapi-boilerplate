package provider

import (
	"github.com/n4de4k/go-webapi-boilerplate/app/accessor"
)

type DataComponent interface {
	GetUserAccessor() accessor.UserAccessor
}
