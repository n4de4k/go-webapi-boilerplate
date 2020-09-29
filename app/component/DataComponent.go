package component

import "github.com/n4de4k/web-api-boilerplate/app/accessor"

type DataComponent interface {
	GetUserAccessor() accessor.UserAccessor
}
