package component

import "github.com/n4de4k/web-api-boilerplate/pkg/accessor"

type DataComponent interface {
	GetUserAccessor() accessor.UserAccessor
}
