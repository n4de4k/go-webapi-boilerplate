package component

import (
	"github.com/n4de4k/web-api-boilerplate/accessor"
)

type DataComponent interface {
	GetUserAccessor() accessor.UserAccessor
}
